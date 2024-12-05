package controller

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	gpuv1alpha1 "iluvatar.ai/ix-gpu-operator/api/v1alpha1"
	"iluvatar.ai/ix-gpu-operator/internal/util"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// singleton: User configuration
// controlFunc: The functions of the service.
// controls: Various services
type GPUClusterApplicationController struct {
	ctx    context.Context
	client client.Client
	scheme *runtime.Scheme

	singleton      *gpuv1alpha1.GPUClusterApplication
	resources      []Resources
	controls       []controlFunc
	componentNames []string

	index     int
	namespace string

	runtime gpuv1alpha1.Runtime
}

func addState(c *GPUClusterApplicationController, path string) (err error) {
	if !util.FileExists(path) {
		err = fmt.Errorf("file does not exist: %v", err)
		return err
	}
	res, cf := addRescourcesControls(path)
	c.resources = append(c.resources, res)
	c.controls = append(c.controls, cf)
	c.componentNames = append(c.componentNames, filepath.Base(path))

	return nil
}

func (c *GPUClusterApplicationController) init(ctx context.Context, reconciler *GPUClusterApplicationReconciler, gpuCluster *gpuv1alpha1.GPUClusterApplication) error {
	c.ctx = ctx
	c.index = 0
	c.client = reconciler.Client
	c.scheme = reconciler.Scheme
	c.singleton = gpuCluster

	logrus.Infof("--- Start %s ---\n", c.singleton.Name)
	if len(c.controls) == 0 {
		gpuClusterCtl.namespace = os.Getenv("OPERATOR_NAMESPACE")
		if gpuClusterCtl.namespace == "" {
			logrus.Errorln("operator namespace environment variable not set")
			os.Exit(1)
		}
		if err := addState(c, "/opt/ix-gpu-operator/ix-device-plugin"); err != nil {
			return err
		}
		if err := addState(c, "/opt/ix-gpu-operator/ix-exporter"); err != nil {
			return err
		}
	}
	logrus.Infoln("assets numbers:", len(c.controls))
	return nil
}

func (c *GPUClusterApplicationController) step() (gpuv1alpha1.State, error) {
	result := gpuv1alpha1.Ready

	logrus.Infof("controls index: %d\n", c.index)
	for _, fs := range c.controls[c.index] {
		stat, err := fs(*c)
		if err != nil {
			return stat, err
		}

		if stat != gpuv1alpha1.Ready {
			result = stat
		}
	}

	// install the next component
	c.index++

	return result, nil
}

func (c *GPUClusterApplicationController) last() bool {
	return c.index == len(c.controls)
}

func (c *GPUClusterApplicationController) isStateEnabled(name string) bool {
	GPUClusterApplicationSpec := &c.singleton.Spec
	switch name {
	case "ix-device-plugin":
		return GPUClusterApplicationSpec.IxDevicePlugin.IsEnabled()
	case "ix-exporter":
		return GPUClusterApplicationSpec.IxExporter.IsEnabled()
	default:
		return false
	}
}
