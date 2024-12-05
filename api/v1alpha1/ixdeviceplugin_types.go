package v1alpha1

import corev1 "k8s.io/api/core/v1"

type IxDevicePluginSpec struct {
	// Enabled indicates whether the ix device plugin is enabled
	Enabled *bool `json:"enabled,omitempty"`

	Repository string `json:"repository,omitempty"`

	Image string `json:"image,omitempty"`

	Version string `json:"version,omitempty"`

	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	ImagePullSecrets []string `json:"imagePullSecrets,omitempty"`

	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	Args []string `json:"args,omitempty"`

	Env []corev1.EnvVar `json:"env,omitempty"`

	Config *ixDevicePluginConfig `json:"config,omitempty"`
}

type ixDevicePluginConfig struct {
	Name    string `json:"name,omitempty"`
	Default string `json:"default,omitempty"`
}

func (idp *IxDevicePluginSpec) IsEnabled() bool {
	if idp.Enabled == nil {
		return true
	}
	return *idp.Enabled
}
