package v1alpha1

import corev1 "k8s.io/api/core/v1"

type IxExporterSpec struct {
	// Enabled indicates whether the ix-exporter is enabled
	Enabled *bool `json:"enabled,omitempty"`

	Repository string `json:"repository,omitempty"`

	Image string `json:"image,omitempty"`

	Version string `json:"version,omitempty"`

	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	ImagePullSecrets []string `json:"imagePullSecrets,omitempty"`

	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	Args []string `json:"args,omitempty"`

	Env []corev1.EnvVar `json:"env,omitempty"`
}

func (e *IxExporterSpec) IsEnabled() bool {
	if e.Enabled == nil {
		return true
	}
	return *e.Enabled
}
