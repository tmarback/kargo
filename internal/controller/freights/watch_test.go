package freights

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	authnv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"

	kargoapi "github.com/akuity/kargo/api/v1alpha1"
	fakeevent "github.com/akuity/kargo/internal/kubernetes/event/fake"
)

func Test_newWatcher(t *testing.T) {
	r := newWatcher(
		&fakeevent.EventRecorder{},
	)
	require.NotNil(t, r.recorder)
}

func Test_watcherUpdate(t *testing.T) {
	testCases := map[string]struct {
		event      event.UpdateEvent
		assertions func(*testing.T, *fakeevent.EventRecorder)
	}{
		"skip if no new approvals": {
			event: event.UpdateEvent{
				ObjectOld: &kargoapi.Freight{
					Status: kargoapi.FreightStatus{
						ApprovedFor: map[string]kargoapi.ApprovedStage{
							"fake-stage": {},
						},
					},
				},
				ObjectNew: &kargoapi.Freight{
					Status: kargoapi.FreightStatus{
						ApprovedFor: map[string]kargoapi.ApprovedStage{
							"fake-stage": {},
						},
					},
				},
			},
			assertions: func(t *testing.T, recorder *fakeevent.EventRecorder) {
				require.Empty(t, recorder.Events)
			},
		},
		"ignore new approvals from the controlplane": {
			event: event.UpdateEvent{
				ObjectOld: &kargoapi.Freight{},
				ObjectNew: &kargoapi.Freight{
					Status: kargoapi.FreightStatus{
						ApprovedFor: map[string]kargoapi.ApprovedStage{
							"fake-stage": {},
						},
					},
				},
			},
			assertions: func(t *testing.T, recorder *fakeevent.EventRecorder) {
				require.Empty(t, recorder.Events)
			},
		},
		"record new approvals": {
			event: event.UpdateEvent{
				ObjectOld: &kargoapi.Freight{},
				ObjectNew: &kargoapi.Freight{
					Status: kargoapi.FreightStatus{
						LastApprovedBy: kargoapi.FormatEventKubernetesUserActor(
							authnv1.UserInfo{
								Username: "fake-user",
							},
						),
						ApprovedFor: map[string]kargoapi.ApprovedStage{
							"fake-stage": {},
						},
					},
				},
			},
			assertions: func(t *testing.T, recorder *fakeevent.EventRecorder) {
				require.Len(t, recorder.Events, 1)

				e := <-recorder.Events
				require.Equal(t, corev1.EventTypeNormal, e.EventType)
				require.Equal(t, kargoapi.EventReasonFreightApproved, e.Reason)
				require.Equal(
					t,
					fmt.Sprintf(
						"Freight approved for Stage %q by %q",
						"fake-stage",
						kargoapi.FormatEventKubernetesUserActor(
							authnv1.UserInfo{
								Username: "fake-user",
							},
						),
					),
					e.Message,
				)
			},
		},
	}
	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			recorder := fakeevent.NewEventRecorder(2)
			w := newWatcher(recorder)
			w.Update(context.TODO(), tc.event, nil)
			tc.assertions(t, recorder)
		})
	}
}
