package controller

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
	"iluvatar.ai/ix-gpu-operator/internal/util"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

type resourcesFromAssets []byte

type Resources struct {
	ServiceAccount     corev1.ServiceAccount
	Role               rbacv1.Role
	ClusterRole        rbacv1.ClusterRole
	RoleBinding        rbacv1.RoleBinding
	ClusterRoleBinding rbacv1.ClusterRoleBinding
	ConfigMaps         []corev1.ConfigMap
	DaemonSet          appsv1.DaemonSet
	Service            corev1.Service
}

func getResouces(path string) []resourcesFromAssets {
	manifests := []resourcesFromAssets{}
	files, err := util.FilePathWithDir(path)
	if err != nil {
		panic(err)
	}
	sort.Strings(files)
	for _, file := range files {
		buffer, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		manifests = append(manifests, resourcesFromAssets(buffer))
	}
	return manifests
}

func addRescourcesControls(path string) (Resources, controlFunc) {
	cf := controlFunc{}
	res := Resources{}

	logrus.Infoln("Get assets from path: ", path)
	manifests := getResouces(path)

	s := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme,
		scheme.Scheme, json.SerializerOptions{Yaml: true, Pretty: false, Strict: false})
	reg := regexp.MustCompile(`\b(\w*kind:\w*)\B.*\b`)

	for _, m := range manifests {
		kind := reg.FindString(string(m))
		sp := strings.Split(kind, ":")
		kind = strings.TrimSpace(sp[1])

		logrus.Infoln("Looking for", "Kind", kind, "in path:", path)
		switch kind {
		case "ServiceAccount":
			_, _, err := s.Decode(m, nil, &res.ServiceAccount)
			if err != nil {
				panic(err)
			}
			cf = append(cf, ServiceAccount)
		case "Role":
			_, _, err := s.Decode(m, nil, &res.Role)
			if err != nil {
				panic(err)
			}
			cf = append(cf, Role)
		case "ClusterRole":
			_, _, err := s.Decode(m, nil, &res.ClusterRole)
			if err != nil {
				panic(err)
			}
			cf = append(cf, ClusterRole)
		case "RoleBinding":
			_, _, err := s.Decode(m, nil, &res.RoleBinding)
			if err != nil {
				panic(err)
			}
			cf = append(cf, RoleBinding)
		case "ClusterRoleBinding":
			_, _, err := s.Decode(m, nil, &res.ClusterRoleBinding)
			if err != nil {
				panic(err)
			}
			cf = append(cf, ClusterRoleBinding)
		case "ConfigMap":
			cm := corev1.ConfigMap{}
			_, _, err := s.Decode(m, nil, &cm)
			if err != nil {
				panic(err)
			}
			res.ConfigMaps = append(res.ConfigMaps, cm)
			if len(res.ConfigMaps) == 1 {
				cf = append(cf, ConfigMaps)
			}
		case "DaemonSet":
			_, _, err := s.Decode(m, nil, &res.DaemonSet)
			if err != nil {
				panic(err)
			}
			cf = append(cf, DaemonSet)
		case "Service":
			_, _, err := s.Decode(m, nil, &res.Service)
			if err != nil {
				panic(err)
			}
			cf = append(cf, Service)
		default:
			fmt.Println("Unknown kind:", kind)
		}
	}
	return res, cf
}
