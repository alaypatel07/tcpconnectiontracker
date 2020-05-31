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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TrackSpec defines the desired state of Track
type TrackSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// PodNamespacedName is the target pod for tracking network connections
	// 1. The input will be in the form of <pod-namespace>/<pod-name>. If "/"
	//    is not found, the namespace of the controller will be used.
	// for pods in different namespace, the controller will RBAC permissions
	// to create, watch, get, list pods in that namespace.
	PodNamespacedName string `json:"podNamespacedName,omitempty"`
}

// TrackStatus defines the observed state of Track
type TrackStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Track is the Schema for the tracks API
type Track struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrackSpec   `json:"spec,omitempty"`
	Status TrackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrackList contains a list of Track
type TrackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Track `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Track{}, &TrackList{})
}
