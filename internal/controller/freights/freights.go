package freights

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	kargoapi "github.com/akuity/kargo/api/v1alpha1"
	"github.com/akuity/kargo/internal/controller"
	libEvent "github.com/akuity/kargo/internal/kubernetes/event"
)

var (
	_ reconcile.Reconciler = &reconciler{}
)

// reconciler reconciles Freight resources.
type reconciler struct{}

// SetupReconcilerWithManager initializes a reconciler for Namespace resources
// and registers it with the provided Manager.
func SetupReconcilerWithManager(
	ctx context.Context,
	kargoMgr manager.Manager,
) error {
	return ctrl.NewControllerManagedBy(kargoMgr).
		Watches(
			&kargoapi.Freight{},
			newWatcher(
				libEvent.NewRecorder(
					ctx,
					kargoMgr.GetScheme(),
					kargoMgr.GetClient(),
					"freight-watcher",
				),
			),
		).
		For(&kargoapi.Freight{}).
		WithOptions(controller.CommonOptions()).
		Complete(newReconciler())
}

func newReconciler() *reconciler {
	return &reconciler{}
}

func (r *reconciler) Reconcile(
	context.Context,
	ctrl.Request,
) (ctrl.Result, error) {
	// No-op
	return ctrl.Result{}, nil
}
