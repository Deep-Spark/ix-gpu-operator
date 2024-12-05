/*
Copyright 2024 corex.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	gpuoperatorv1alpha1 "iluvatar.ai/ix-gpu-operator/api/v1alpha1"
)

var gpuClusterCtl GPUClusterApplicationController

// GPUClusterApplicationReconciler reconciles a GPUClusterApplication object
type GPUClusterApplicationReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=iluvatar.com,resources=gpuclusterapplications,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=iluvatar.com,resources=gpuclusterapplications/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=iluvatar.com,resources=gpuclusterapplications/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=namespaces;serviceaccounts;pods;pods/eviction;services;services/finalizers;endpoints,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the GPUClusterApplication object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *GPUClusterApplicationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	fmt.Println("--- Reconcile Func ---")
	gpuObjects := gpuoperatorv1alpha1.GPUClusterApplication{}
	err := r.Client.Get(ctx, req.NamespacedName, &gpuObjects)
	if err != nil {
		// if return result is not nil, the request will be requeued.
		// if return result is nil, the request will be ignored.
		if client.IgnoreNotFound(err) != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	// the object has been marked for deletion and not to requeued.
	if gpuObjects.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}
	err = gpuClusterCtl.init(ctx, r, &gpuObjects)
	if err != nil {
		err = fmt.Errorf("failed to init gpuclusterctl: %v", err)
		return ctrl.Result{}, err
	}

	overallStatus := gpuoperatorv1alpha1.Ready

	for {
		// Requeue and deployment
		state, err := gpuClusterCtl.step()
		if err != nil {
			logrus.Errorf("Deployment error: %v", err)
			return ctrl.Result{
				RequeueAfter: time.Second * 15,
			}, err
		}

		if state == gpuoperatorv1alpha1.NotReady {
			logrus.Warningf("Assets state: %v\n", state)
			overallStatus = gpuoperatorv1alpha1.NotReady
		}

		if gpuClusterCtl.last() {
			break
		}
	}

	if overallStatus != gpuoperatorv1alpha1.Ready {
		gpuObjects.Status.State = overallStatus
		logrus.Infoln("Assets Is NotReady.")
		return ctrl.Result{
			RequeueAfter: time.Second * 2,
		}, nil
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GPUClusterApplicationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&gpuoperatorv1alpha1.GPUClusterApplication{}).
		Complete(r)
}
