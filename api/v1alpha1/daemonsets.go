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
	corev1 "k8s.io/api/core/v1"
)

// OperatorSpec describes configuration options for the operator
type DaemonsetsSpec struct {
	Labels map[string]string `json:"labels,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`

	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	RollingUpdate *RollingUpdateSpec `json:"rollingUpdate,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`

	PriorityClassName string `json:"priorityClassName,omitempty"`
}

type RollingUpdateSpec struct {
	MaxUnavailable string `json:"maxUnavailable,omitempty"`
}
