package controller

import (
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	api "github.com/akuity/kargo/api/v1alpha1"
)

func TestGetShardPredicate(t *testing.T) {
	const testShardName = "test-shard"
	unlabeledEvent := event.CreateEvent{
		Object: &api.Stage{},
	}
	labeledEvent := event.CreateEvent{
		Object: &api.Stage{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					ShardLabelKey: testShardName,
				},
			},
		},
	}
	testCases := []struct {
		name       string
		shardName  string
		assertions func(predicate.Predicate, error)
	}{
		{
			name:      "shard name is the empty string",
			shardName: "",
			assertions: func(pred predicate.Predicate, err error) {
				require.NoError(t, err)
				require.NotNil(t, pred)
				require.True(t, pred.Create(unlabeledEvent))
				require.False(t, pred.Create(labeledEvent))
			},
		},
		{
			name:      "shard name is not the empty string",
			shardName: testShardName,
			assertions: func(pred predicate.Predicate, err error) {
				require.NoError(t, err)
				require.NotNil(t, pred)
				require.False(t, pred.Create(unlabeledEvent))
				require.True(t, pred.Create(labeledEvent))
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.assertions(GetShardPredicate(testCase.shardName))
		})
	}
}