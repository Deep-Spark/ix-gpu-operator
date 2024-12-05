package v1alpha1

import (
	"fmt"
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

func imagePath(repository, image, version, imagePathEnvName string) (string, error) {
	var crdImagePath string
	if repository == "" && version == "" {
		if image != "" {
			crdImagePath = image
		}
	} else {
		if strings.HasPrefix(version, "sha256:") {
			crdImagePath = repository + "/" + image + "@" + version
		} else {
			crdImagePath = repository + "/" + image + ":" + version
		}
	}
	if crdImagePath != "" {
		return crdImagePath, nil
	}

	envImagePath := os.Getenv(imagePathEnvName)
	if envImagePath != "" {
		return envImagePath, nil
	}

	return "", fmt.Errorf("empty image path: %s", imagePathEnvName)
}

func ImagePath(spec interface{}) (string, error) {
	switch v := spec.(type) {
	case *IxDevicePluginSpec:
		config := spec.(*IxDevicePluginSpec)
		return imagePath(config.Repository, config.Image, config.Version, "IX_DEVICE_PLUGIN")
	case *IxExporterSpec:
		config := spec.(*IxExporterSpec)
		return imagePath(config.Repository, config.Image, config.Version, "IX_EXPORTER")
	default:
		return "", fmt.Errorf("invalid type to construct image type path: %v", v)
	}
}

func ImagePullPolicy(pullPolicy string) (imagePullPolicy corev1.PullPolicy) {
	switch pullPolicy {
	case "Always":
		imagePullPolicy = corev1.PullAlways
	case "IfNotPresent":
		imagePullPolicy = corev1.PullIfNotPresent
	case "Never":
		imagePullPolicy = corev1.PullNever
	default:
		imagePullPolicy = corev1.PullIfNotPresent

	}
	return imagePullPolicy
}
