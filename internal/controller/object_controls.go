package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	gpuv1alpha1 "iluvatar.ai/ix-gpu-operator/api/v1alpha1"
	"iluvatar.ai/ix-gpu-operator/internal/util"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	controllerutil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type controlFunc []func(c GPUClusterApplicationController) (gpuv1alpha1.State, error)

func ServiceAccount(c GPUClusterApplicationController) (gpuv1alpha1.State, error) {
	ctx := c.ctx
	index := c.index
	saObj := c.resources[index].ServiceAccount.DeepCopy()
	saObj.Namespace = c.namespace

	logrus.Infof("%s: ServiceAccount", saObj.Name)
	// Check if the component is enabled or not.
	if !c.isStateEnabled(c.componentNames[index]) {
		err := c.client.Delete(ctx, saObj)
		if err != nil && !apierrors.IsNotFound(err) {
			logrus.Infof("Failed to delete ServiceAccount: %s\n", saObj.Name)
			return gpuv1alpha1.NotReady, err
		}
		return gpuv1alpha1.Disabled, nil
	}
	// Associate resources with controllers
	if err := controllerutil.SetControllerReference(c.singleton, saObj, c.scheme); err != nil {
		return gpuv1alpha1.NotReady, err
	}
	if err := c.client.Create(ctx, saObj); err != nil {
		if apierrors.IsAlreadyExists(err) {
			logrus.Infoln("Found Resource, update")
			err = c.client.Update(ctx, saObj)
			if err != nil {
				logrus.Infof("Failed to update ServiceAccount: %s\n", saObj.Name)
				return gpuv1alpha1.NotReady, err
			}
			return gpuv1alpha1.Ready, nil
		}
		logrus.Infof("Failed to create ServiceAccount: %s\n", saObj.Name)
		return gpuv1alpha1.NotReady, err
	}
	return gpuv1alpha1.Ready, nil
}

func Role(c GPUClusterApplicationController) (gpuv1alpha1.State, error) {
	ctx := c.ctx
	index := c.index
	roleObj := c.resources[index].Role.DeepCopy()
	roleObj.Namespace = c.namespace

	logrus.Infof("%s: role", roleObj.Name)
	if !c.isStateEnabled(c.componentNames[index]) {
		err := c.client.Delete(ctx, roleObj)
		if err != nil && !apierrors.IsNotFound(err) {
			logrus.Infof("Failed to delete role: %s\n", roleObj.Name)
			return gpuv1alpha1.NotReady, err
		}
		return gpuv1alpha1.Disabled, nil
	}
	// Associate resources with controllers
	if err := controllerutil.SetControllerReference(c.singleton, roleObj, c.scheme); err != nil {
		return gpuv1alpha1.NotReady, err
	}

	if err := c.client.Create(ctx, roleObj); err != nil {
		if apierrors.IsAlreadyExists(err) {
			logrus.Infoln("Found Resource, update")
			err = c.client.Update(ctx, roleObj)
			if err != nil {
				logrus.Infof("Failed to update Role: %s\n", roleObj.Name)
				return gpuv1alpha1.NotReady, err
			}
			return gpuv1alpha1.Ready, nil
		}
		logrus.Infof("Failed to create Role: %s\n", roleObj.Name)
		return gpuv1alpha1.NotReady, err
	}
	return gpuv1alpha1.Ready, nil
}

func ClusterRole(c GPUClusterApplicationController) (gpuv1alpha1.State, error) {
	ctx := c.ctx
	index := c.index
	clusterRoleObj := c.resources[index].ClusterRole.DeepCopy()
	clusterRoleObj.Namespace = c.namespace

	logrus.Infof("%s: clusterRole", clusterRoleObj.Name)
	// Check if the component is enabled or not.
	if !c.isStateEnabled(c.componentNames[index]) {
		err := c.client.Delete(ctx, clusterRoleObj)
		if err != nil && !apierrors.IsNotFound(err) {
			logrus.Infof("Failed to delete clusterRole: %s\n", clusterRoleObj.Name)
			return gpuv1alpha1.NotReady, err
		}
		return gpuv1alpha1.Disabled, nil
	}
	// Associate resources with controllers
	if err := controllerutil.SetControllerReference(c.singleton, clusterRoleObj, c.scheme); err != nil {
		return gpuv1alpha1.NotReady, err
	}

	if err := c.client.Create(ctx, clusterRoleObj); err != nil {
		if apierrors.IsAlreadyExists(err) {
			logrus.Infoln("Found Resource, update")
			err = c.client.Update(ctx, clusterRoleObj)
			if err != nil {
				logrus.Infof("Failed to update clusterRole: %s\n", clusterRoleObj.Name)
				return gpuv1alpha1.NotReady, err
			}
			return gpuv1alpha1.Ready, nil
		}
		logrus.Infof("Failed to create clusterRole: %s\n", clusterRoleObj.Name)
		return gpuv1alpha1.NotReady, err
	}
	return gpuv1alpha1.Ready, nil
}

func RoleBinding(c GPUClusterApplicationController) (gpuv1alpha1.State, error) {
	ctx := c.ctx
	index := c.index
	roleBindingObj := c.resources[index].RoleBinding.DeepCopy()
	roleBindingObj.Namespace = c.namespace

	logrus.Infof("%s: RoleBinding", roleBindingObj.Name)
	// Check if the component is enabled or not.
	if !c.isStateEnabled(c.componentNames[index]) {
		err := c.client.Delete(ctx, roleBindingObj)
		if err != nil && !apierrors.IsNotFound(err) {
			logrus.Infof("Failed to delete roleBinding: %s\n", roleBindingObj.Name)
			return gpuv1alpha1.NotReady, err
		}
		return gpuv1alpha1.Disabled, nil
	}
	// Associate resources with controllers
	if err := controllerutil.SetControllerReference(c.singleton, roleBindingObj, c.scheme); err != nil {
		return gpuv1alpha1.NotReady, err
	}

	if err := c.client.Create(ctx, roleBindingObj); err != nil {
		if apierrors.IsAlreadyExists(err) {
			logrus.Infoln("Found Resource, update")
			err = c.client.Update(ctx, roleBindingObj)
			if err != nil {
				logrus.Infof("Failed to update roleBinding: %s\n", roleBindingObj.Name)
				return gpuv1alpha1.NotReady, err
			}
			return gpuv1alpha1.Ready, nil
		}
		logrus.Infof("Failed to create roleBinding: %s\n", roleBindingObj.Name)
		return gpuv1alpha1.NotReady, err
	}
	return gpuv1alpha1.Ready, nil
}

func ClusterRoleBinding(c GPUClusterApplicationController) (gpuv1alpha1.State, error) {
	ctx := c.ctx
	index := c.index
	clusterRoleBindingObj := c.resources[index].ClusterRoleBinding.DeepCopy()
	clusterRoleBindingObj.Namespace = c.namespace

	logrus.Infof("%s: clusterRoleBinding", clusterRoleBindingObj.Name)
	// Check if the component is enabled or not.
	if !c.isStateEnabled(c.componentNames[index]) {
		err := c.client.Delete(ctx, clusterRoleBindingObj)
		if err != nil && !apierrors.IsNotFound(err) {
			logrus.Infof("Failed to delete clusterRoleBinding: %s\n", clusterRoleBindingObj.Name)
			return gpuv1alpha1.NotReady, err
		}
		return gpuv1alpha1.Disabled, nil
	}
	// Associate resources with controllers
	if err := controllerutil.SetControllerReference(c.singleton, clusterRoleBindingObj, c.scheme); err != nil {
		return gpuv1alpha1.NotReady, err
	}

	if err := c.client.Create(ctx, clusterRoleBindingObj); err != nil {
		if apierrors.IsAlreadyExists(err) {
			logrus.Infoln("Found Resource, update")
			err = c.client.Update(ctx, clusterRoleBindingObj)
			if err != nil {
				logrus.Infof("Failed to update clusterRoleBinding: %s\n", clusterRoleBindingObj.Name)
				return gpuv1alpha1.NotReady, err
			}
			return gpuv1alpha1.Ready, nil
		}
		logrus.Infof("Failed to create clusterRoleBinding: %s\n", clusterRoleBindingObj.Name)
		return gpuv1alpha1.NotReady, err
	}
	return gpuv1alpha1.Ready, nil
}

func createConfigMap(c GPUClusterApplicationController, cmIdx int) (gpuv1alpha1.State, error) {
	ctx := c.ctx
	index := c.index
	cmObj := c.resources[index].ConfigMaps[cmIdx].DeepCopy()
	cmObj.Namespace = c.namespace

	logrus.Infof("%s: configMap", cmObj.Name)
	// Check if the component is enabled or not.
	if !c.isStateEnabled(c.componentNames[index]) {
		err := c.client.Delete(ctx, cmObj)
		if err != nil && !apierrors.IsNotFound(err) {
			logrus.Infof("Failed to delete configmap: %s\n", cmObj.Name)
			return gpuv1alpha1.NotReady, err
		}
		return gpuv1alpha1.Disabled, nil
	}

	// Associate resources with controllers
	if err := controllerutil.SetOwnerReference(c.singleton, cmObj, c.scheme); err != nil {
		return gpuv1alpha1.NotReady, err
	}

	if err := c.client.Create(ctx, cmObj); err != nil {
		if apierrors.IsAlreadyExists(err) {
			logrus.Infoln("Found Resource, update")
			err = c.client.Update(ctx, cmObj)
			if err != nil {
				logrus.Infof("Failed to update configmap: %s\n", cmObj.Name)
				return gpuv1alpha1.NotReady, err
			}
			return gpuv1alpha1.Ready, nil
		}
		logrus.Infof("Failed to create configmap: %s\n", cmObj.Name)
		return gpuv1alpha1.NotReady, err
	}
	return gpuv1alpha1.Ready, nil
}

func ConfigMaps(c GPUClusterApplicationController) (gpuv1alpha1.State, error) {
	index := c.index
	status := gpuv1alpha1.Ready
	for i := range c.resources[index].ConfigMaps {
		status, err := createConfigMap(c, i)
		if err != nil {
			return status, err
		}

		if status == gpuv1alpha1.NotReady {
			status = gpuv1alpha1.NotReady
		}
	}
	return status, nil
}

func TransformIxDevicePlugin(ds *appsv1.DaemonSet, config *gpuv1alpha1.GPUClusterApplicationSpec, c GPUClusterApplicationController) error {
	image, err := gpuv1alpha1.ImagePath(&config.IxDevicePlugin)
	if err != nil {
		return err
	}
	ds.Spec.Template.Spec.Containers[0].Image = image

	// update image pull policy
	ds.Spec.Template.Spec.Containers[0].ImagePullPolicy = gpuv1alpha1.ImagePullPolicy(config.IxDevicePlugin.ImagePullPolicy)

	// set image pull secrets
	if len(config.IxDevicePlugin.ImagePullSecrets) > 0 {
		for _, secret := range config.IxDevicePlugin.ImagePullSecrets {
			if !util.SetContainerSecret(ds.Spec.Template.Spec.ImagePullSecrets, secret) {
				ds.Spec.Template.Spec.ImagePullSecrets = append(ds.Spec.Template.Spec.ImagePullSecrets, corev1.LocalObjectReference{
					Name: secret,
				})
			}
		}
	}

	// set arguments if specified for ix-device-plugin container
	if len(config.IxDevicePlugin.Args) > 0 {
		ds.Spec.Template.Spec.Containers[0].Args = config.IxDevicePlugin.Args
	}

	// set environment if specified for ix-device-plugin container
	if len(config.IxDevicePlugin.Env) > 0 {
		for _, env := range config.IxDevicePlugin.Env {
			util.SetContainerEnv(&ds.Spec.Template.Spec.Containers[0], env.Name, env.Value)
		}
	}

	// set resources limit
	if config.IxDevicePlugin.Resources != nil {
		for i := range ds.Spec.Template.Spec.Containers {
			ds.Spec.Template.Spec.Containers[i].Resources.Requests = config.IxDevicePlugin.Resources.Requests
			ds.Spec.Template.Spec.Containers[i].Resources.Limits = config.IxDevicePlugin.Resources.Limits
		}
	}

	return nil
}

func TransformIxExporter(ds *appsv1.DaemonSet, config *gpuv1alpha1.GPUClusterApplicationSpec, c GPUClusterApplicationController) error {
	image, err := gpuv1alpha1.ImagePath(&config.IxExporter)
	if err != nil {
		return err
	}
	ds.Spec.Template.Spec.Containers[0].Image = image

	// update image pull policy
	ds.Spec.Template.Spec.Containers[0].ImagePullPolicy = gpuv1alpha1.ImagePullPolicy(config.IxExporter.ImagePullPolicy)

	// set image pull secrets
	if len(config.IxExporter.ImagePullSecrets) > 0 {
		for _, secret := range config.IxExporter.ImagePullSecrets {
			if !util.SetContainerSecret(ds.Spec.Template.Spec.ImagePullSecrets, secret) {
				ds.Spec.Template.Spec.ImagePullSecrets = append(ds.Spec.Template.Spec.ImagePullSecrets, corev1.LocalObjectReference{
					Name: secret,
				})
			}
		}
	}

	// set arguments if specified for ix-exporter container
	if len(config.IxExporter.Args) > 0 {
		ds.Spec.Template.Spec.Containers[0].Args = config.IxExporter.Args
	}

	// set environment if specified for ix-exporter container
	if len(config.IxExporter.Env) > 0 {
		for _, env := range config.IxExporter.Env {
			util.SetContainerEnv(&ds.Spec.Template.Spec.Containers[0], env.Name, env.Value)
		}
	}

	// set resources limit
	if config.IxExporter.Resources != nil {
		for i := range ds.Spec.Template.Spec.Containers {
			ds.Spec.Template.Spec.Containers[i].Resources.Requests = config.IxExporter.Resources.Requests
			ds.Spec.Template.Spec.Containers[i].Resources.Limits = config.IxExporter.Resources.Limits
		}
	}

	return nil
}

func applyCommonDaemonSetConfig(ds *appsv1.DaemonSet, config *gpuv1alpha1.GPUClusterApplicationSpec) error {
	switch config.Daemonsets.UpdateStrategy {
	case "OnDelete":
		ds.Spec.UpdateStrategy = appsv1.DaemonSetUpdateStrategy{
			Type: appsv1.OnDeleteDaemonSetStrategyType,
		}
	case "RollingUpdate":
		fallthrough
	default:
		if config.Daemonsets.RollingUpdate == nil || config.Daemonsets.RollingUpdate.MaxUnavailable == "" {
			return nil
		}
		var intOrString intstr.IntOrString
		if strings.HasPrefix(config.Daemonsets.RollingUpdate.MaxUnavailable, "%") {
			intOrString = intstr.IntOrString{
				Type:   intstr.String,
				StrVal: config.Daemonsets.RollingUpdate.MaxUnavailable,
			}
		} else {
			int64Val, err := strconv.ParseInt(config.Daemonsets.RollingUpdate.MaxUnavailable, 10, 32)
			if err != nil {
				return fmt.Errorf("failed to parse maxUnavailable for RollingUpdate: %v", err)
			}
			intOrString = intstr.IntOrString{
				Type:   intstr.Int,
				IntVal: int32(int64Val),
			}
		}
		rollingUpdateSpec := appsv1.RollingUpdateDaemonSet{
			MaxUnavailable: &intOrString,
		}
		ds.Spec.UpdateStrategy = appsv1.DaemonSetUpdateStrategy{
			Type:          appsv1.RollingUpdateDaemonSetStrategyType,
			RollingUpdate: &rollingUpdateSpec,
		}

		// set PriorityClassName if specified
		if config.Daemonsets.PriorityClassName != "" {
			ds.Spec.Template.Spec.PriorityClassName = config.Daemonsets.PriorityClassName
		}

		if len(config.Daemonsets.Tolerations) > 0 {
			ds.Spec.Template.Spec.Tolerations = config.Daemonsets.Tolerations
		}
	}

	return nil
}

func configDaemonSetMetaData(ds *appsv1.DaemonSet, configDsSpec *gpuv1alpha1.DaemonsetsSpec) {
	if len(configDsSpec.Labels) > 0 {
		if ds.Spec.Template.ObjectMeta.Labels == nil {
			ds.Spec.Template.ObjectMeta.Labels = make(map[string]string)
		}
		for key, value := range configDsSpec.Labels {
			if key == "app" || key == "app.kubernetes.io/part-of" {
				continue
			}
			ds.Spec.Template.ObjectMeta.Labels[key] = value
		}
	}

	if len(configDsSpec.Annotations) > 0 {
		if ds.Spec.Template.ObjectMeta.Annotations == nil {
			ds.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		for key, value := range configDsSpec.Annotations {
			ds.Spec.Template.ObjectMeta.Annotations[key] = value
		}
	}
}

// pre-config for DaemonSet: fillful daemonset with configuration file
func preConfigDaemonSet(c GPUClusterApplicationController, daemonSetObj *appsv1.DaemonSet) error {
	transformations := map[string]func(*appsv1.DaemonSet, *gpuv1alpha1.GPUClusterApplicationSpec, GPUClusterApplicationController) error{
		"iluvatar-device-plugin": TransformIxDevicePlugin,
		"iluvatar-ix-exporter":   TransformIxExporter,
	}
	fs, ok := transformations[daemonSetObj.Name]
	if !ok {
		fmt.Printf("No transformation for %s\n", daemonSetObj.Name)
		return nil
	}

	err := applyCommonDaemonSetConfig(daemonSetObj, &c.singleton.Spec)
	if err != nil {
		return fmt.Errorf("failed to apply common daemonSet config: %s", daemonSetObj.Name)
	}

	err = fs(daemonSetObj, &c.singleton.Spec, c)
	if err != nil {
		return fmt.Errorf("failed to apply %s transformation: %s", daemonSetObj.Name, err)
	}

	configDaemonSetMetaData(daemonSetObj, &c.singleton.Spec.Daemonsets)
	return nil
}

func getPodsOwnedByDaemonSet(ds *appsv1.DaemonSet, pods []corev1.Pod, c GPUClusterApplicationController) []corev1.Pod {
	dsPodList := []corev1.Pod{}
	for _, pod := range pods {
		if pod.OwnerReferences == nil || len(pod.OwnerReferences) < 1 {
			logrus.Infof("pod %s has no owner DaemonSet\n", pod.Name)
			continue
		}
		if ds.UID != pod.OwnerReferences[0].UID {
			logrus.Infof("pod %s is not owned by %s\n", pod.Name, ds.Name)
			continue
		}
		dsPodList = append(dsPodList, pod)
	}
	return dsPodList
}

func checkDaemonsetReady(name string, c GPUClusterApplicationController) gpuv1alpha1.State {
	ctx := c.ctx
	ds := &appsv1.DaemonSet{}
	err := c.client.Get(ctx, types.NamespacedName{
		Name:      name,
		Namespace: c.namespace,
	}, ds)
	if err != nil {
		logrus.Infof("failed to get daemonset: %s,%v\n", name, err)
		return gpuv1alpha1.NotReady
	}

	// The number of pods expected to be deployed by daemonset
	if ds.Status.DesiredNumberScheduled == 0 {
		return gpuv1alpha1.Ready
	}

	if ds.Status.NumberUnavailable != 0 {
		return gpuv1alpha1.NotReady
	}

	if ds.Spec.UpdateStrategy.Type != appsv1.OnDeleteDaemonSetStrategyType {
		return gpuv1alpha1.Ready
	}

	opts := []client.ListOption{client.MatchingLabels(ds.Spec.Template.ObjectMeta.Labels)}
	list := &corev1.PodList{}
	err = c.client.List(ctx, list, opts...)
	if err != nil {
		fmt.Printf("failed to list pods: %s\n", err)
		return gpuv1alpha1.NotReady
	}
	if len(list.Items) == 0 {
		return gpuv1alpha1.NotReady
	}

	dsPods := getPodsOwnedByDaemonSet(ds, list.Items, c)
	daemonsetRevision, err := util.GetDaemonsetControllerRevisionHash(ctx, ds, c.client)
	if err != nil {
		fmt.Printf("failed to get daemonset controller revision hash: %v\n", err)
		return gpuv1alpha1.NotReady
	}

	for _, pod := range dsPods {
		podRevisionHash, err := util.GetPodControllerRevisionHash(&pod)
		if err != nil {
			logrus.Infof("failed to get pod controller revision hash: %v\n", err)
			return gpuv1alpha1.NotReady
		}

		if podRevisionHash != daemonsetRevision || pod.Status.Phase != corev1.PodRunning {
			return gpuv1alpha1.NotReady
		}

		// If the pod generation matches the daemonset generation and the pod is running
		// and it has at least 1 container
		if len(pod.Status.ContainerStatuses) != 0 {
			for i := range pod.Status.ContainerStatuses {
				if !pod.Status.ContainerStatuses[i].Ready {
					return gpuv1alpha1.NotReady
				}
			}
		}
	}

	return gpuv1alpha1.Ready
}

func checkDaemonSetChanged(currentDs *appsv1.DaemonSet, newDs *appsv1.DaemonSet) bool {
	if currentDs == nil && newDs == nil {
		return true
	}
	if currentDs.Annotations == nil || newDs.Annotations == nil {
		panic("appsv1.DaemonSet.Annotations must be allocated prior to calling isDaemonsetSpecChanged()")
	}
	foundHashAnnotation := false
	hashStr := util.GetDaemonsetHash(newDs)
	for key, value := range currentDs.Annotations {
		if key == util.AnnotationsHashKey {
			if value != hashStr {
				// update annotation to daemonset as per new spec
				newDs.Annotations[key] = hashStr
				return true
			}
			foundHashAnnotation = true
			break
		}
	}
	if !foundHashAnnotation {
		newDs.Annotations[util.AnnotationsHashKey] = hashStr
		return true
	}
	return false
}

func DaemonSet(c GPUClusterApplicationController) (gpuv1alpha1.State, error) {
	ctx := c.ctx
	index := c.index
	dsObj := c.resources[index].DaemonSet.DeepCopy()
	dsObj.Namespace = c.namespace

	logrus.Infof("%s: DaemonSet", dsObj.Name)
	if !c.isStateEnabled(c.componentNames[index]) {
		err := c.client.Delete(ctx, dsObj)
		if err != nil && !apierrors.IsNotFound(err) {
			logrus.Infof("Failed to delete daemonSet: %s\n", dsObj.Name)
			return gpuv1alpha1.NotReady, err
		}
		return gpuv1alpha1.Disabled, nil
	}

	err := preConfigDaemonSet(c, dsObj)
	if err != nil {
		return gpuv1alpha1.NotReady, err
	}

	if err := controllerutil.SetControllerReference(c.singleton, dsObj, c.scheme); err != nil {
		return gpuv1alpha1.NotReady, err
	}

	if dsObj.Labels == nil {
		dsObj.Labels = make(map[string]string)
	}
	for key, value := range c.singleton.Spec.Daemonsets.Labels {
		dsObj.Labels[key] = value
	}

	if dsObj.Annotations == nil {
		dsObj.Annotations = make(map[string]string)
	}
	for key, value := range c.singleton.Spec.Daemonsets.Annotations {
		dsObj.Annotations[key] = value
	}

	foundDs := &appsv1.DaemonSet{}
	err = c.client.Get(ctx, types.NamespacedName{
		Name:      dsObj.Name,
		Namespace: dsObj.Namespace,
	}, foundDs)
	if err != nil && apierrors.IsNotFound(err) {
		logrus.Infoln("DaemonSet not found, creating...")
		// generate hash for the daemonset
		hasher := util.GetDaemonsetHash(dsObj)
		// add annotation for the daemonset with hash
		dsObj.Annotations[util.AnnotationsHashKey] = hasher
		err = c.client.Create(ctx, dsObj)
		if err != nil {
			logrus.Infof("Failed to create daemonSet: %s\n", dsObj.Name)
			return gpuv1alpha1.NotReady, err
		}
		return checkDaemonsetReady(dsObj.Name, c), nil
	} else if err != nil {
		logrus.Infof("Failed to get daemonSet: %s\n", dsObj.Name)
		return gpuv1alpha1.NotReady, err
	}

	changed := checkDaemonSetChanged(foundDs, dsObj)
	if changed {
		logrus.Infoln("DaemonSet spec changed, updating...")
		err = c.client.Update(ctx, dsObj)
		if err != nil {
			return gpuv1alpha1.NotReady, err
		}
	}
	return checkDaemonsetReady(dsObj.Name, c), nil
}

func Service(c GPUClusterApplicationController) (gpuv1alpha1.State, error) {
	ctx := c.ctx
	index := c.index
	svcObj := c.resources[index].Service.DeepCopy()
	svcObj.Namespace = c.namespace

	logrus.Infof("%s: Service", svcObj.Name)
	if !c.isStateEnabled(c.componentNames[index]) {
		err := c.client.Delete(ctx, svcObj)
		if err != nil && !apierrors.IsNotFound(err) {
			logrus.Infof("Failed to delete Service: %s\n", svcObj.Name)
			return gpuv1alpha1.NotReady, err
		}
		return gpuv1alpha1.Disabled, nil
	}

	if err := controllerutil.SetControllerReference(c.singleton, svcObj, c.scheme); err != nil {
		return gpuv1alpha1.NotReady, err
	}

	found := &corev1.Service{}
	err := c.client.Get(ctx, types.NamespacedName{
		Namespace: svcObj.Namespace,
		Name:      svcObj.Name,
	}, found)
	if err != nil && apierrors.IsNotFound(err) {
		logrus.Infoln("Service not found, creating...")
		if err := c.client.Create(ctx, svcObj); err != nil {
			logrus.Infof("Failed to create Service: %s\n", svcObj.Name)
			return gpuv1alpha1.NotReady, err
		}
		return gpuv1alpha1.Ready, err
	} else if err != nil {
		return gpuv1alpha1.NotReady, err
	}

	logrus.Infoln("Found Service Resource, updating...")
	svcObj.ResourceVersion = found.ResourceVersion
	svcObj.Spec.Ports = found.Spec.Ports
	return gpuv1alpha1.Ready, nil
}
