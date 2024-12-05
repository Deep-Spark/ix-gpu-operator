package util

import (
	corev1 "k8s.io/api/core/v1"
)

func SetContainerSecret(secrets []corev1.LocalObjectReference, secret string) bool {
	for _, s := range secrets {
		if s.Name == secret {
			return true
		}
	}
	return false
}

func SetContainerEnv(c *corev1.Container, key, value string) {
	for i, v := range c.Env {
		if v.Name != key {
			continue
		}
		c.Env[i].Value = value
		return
	}
	c.Env = append(c.Env, corev1.EnvVar{
		Name:  key,
		Value: value,
	})
}
