package freights

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"

	kargoapi "github.com/akuity/kargo/api/v1alpha1"
)

var (
	_ handler.EventHandler = &watcher{}
)

// watcher watches Freight resources.
type watcher struct {
	recorder record.EventRecorder
}

func newWatcher(
	recorder record.EventRecorder,
) *watcher {
	return &watcher{
		recorder: recorder,
	}
}

func (w *watcher) Create(
	context.Context,
	event.CreateEvent,
	workqueue.RateLimitingInterface,
) {
	// No-op
}

func (w *watcher) Update(
	_ context.Context,
	e event.UpdateEvent,
	_ workqueue.RateLimitingInterface,
) {
	freight := e.ObjectNew.(*kargoapi.Freight)    // nolint: forcetypeassert
	oldFreight := e.ObjectOld.(*kargoapi.Freight) // nolint: forcetypeassert

	if freight.Status.LastApprovedBy != "" {
		for approvedStage := range freight.Status.ApprovedFor {
			if _, ok := oldFreight.Status.ApprovedFor[approvedStage]; !ok {
				w.recorder.AnnotatedEventf(
					freight,
					kargoapi.NewFreightApprovedEventAnnotations(
						freight.Status.LastApprovedBy,
						freight,
						approvedStage,
					),
					corev1.EventTypeNormal,
					kargoapi.EventReasonFreightApproved,
					"Freight approved for Stage %q by %q",
					approvedStage,
					freight.Status.LastApprovedBy,
				)
			}
		}
	}
}

func (w *watcher) Delete(
	context.Context,
	event.DeleteEvent,
	workqueue.RateLimitingInterface,
) {
	// No-op
}

func (w *watcher) Generic(
	context.Context,
	event.GenericEvent,
	workqueue.RateLimitingInterface,
) {
	// No-op
}
