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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GPUClusterApplicationSpec defines the desired state of GPUClusterApplication
type GPUClusterApplicationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Cluster defines configurations for cluster
	Cluster ClusterSpec `json:"cluster"`

	// DaemonSets defines configurations for daemonsets
	Daemonsets DaemonsetsSpec `json:"daemonsets"`

	// ix device plugin component spec
	IxDevicePlugin IxDevicePluginSpec `json:"ixDevicePlugin,omitempty"`

	// ix exporter component spec
	IxExporter IxExporterSpec `json:"ixExporter,omitempty"`
}

// GPUClusterApplicationStatus defines the observed state of GPUClusterApplication
type GPUClusterApplicationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Namespace string `json:"namespace,omitempty"`

	State State `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GPUClusterApplication is the Schema for the gpuclusterapplications API
type GPUClusterApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GPUClusterApplicationSpec   `json:"spec,omitempty"`
	Status GPUClusterApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GPUClusterApplicationList contains a list of GPUClusterApplication
type GPUClusterApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GPUClusterApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GPUClusterApplication{}, &GPUClusterApplicationList{})
}
