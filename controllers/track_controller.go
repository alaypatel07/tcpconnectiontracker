/*


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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tcpnetworkconnectionsiov1alpha1 "github.com/alaypatel07/tcpconnectiontracker/api/v1alpha1"
)

// TrackReconciler reconciles a Track object
type TrackReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tcp.network-connections.io.tcp.network-connections.io,resources=tracks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tcp.network-connections.io.tcp.network-connections.io,resources=tracks/status,verbs=get;update;patch

func (r *TrackReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("track", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *TrackReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tcpnetworkconnectionsiov1alpha1.Track{}).
		Complete(r)
}
