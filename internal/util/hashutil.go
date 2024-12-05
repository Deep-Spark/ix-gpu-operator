package util

import (
	"context"
	"fmt"
	"hash/fnv"
	"sort"
	"strings"

	"github.com/davecgh/go-spew/spew"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	AnnotationsHashKey = "iluvatar.com/annotation-hasher"

	PodControllerRevisionHashLabelKey = "controller-revision-hash"
)

func GetDaemonsetHash(daemonSetObj *appsv1.DaemonSet) string {
	hasher := fnv.New32a()
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}
	printer.Fprintf(hasher, "%#v", daemonSetObj)
	return fmt.Sprint(hasher.Sum32())
}

func GetDaemonsetControllerRevisionHash(ctx context.Context, ds *appsv1.DaemonSet, cli client.Client) (string, error) {
	opts := []client.ListOption{
		client.InNamespace(ds.Namespace),
		client.MatchingLabels(ds.Spec.Selector.MatchLabels),
	}
	list := &appsv1.ControllerRevisionList{}
	if err := cli.List(ctx, list, opts...); err != nil {
		return "", fmt.Errorf("error getting controller revision list for daemonset %s: %v", ds.Name, err)
	}

	var revisions []appsv1.ControllerRevision
	for _, cr := range list.Items {
		if strings.HasPrefix(cr.Name, ds.Name) {
			revisions = append(revisions, cr)
		}
	}
	if len(revisions) == 0 {
		return "", fmt.Errorf("no controller revisions found for daemonset %s", ds.Name)
	}

	sort.Slice(revisions, func(i, j int) bool {
		return revisions[i].Name < revisions[j].Name
	})

	currentRevisions := revisions[len(revisions)-1]
	hash := strings.TrimPrefix(currentRevisions.Name, fmt.Sprintf("%s-", ds.Name))
	return hash, nil
}

func GetPodControllerRevisionHash(pod *corev1.Pod) (string, error) {
	if hash, ok := pod.Labels[PodControllerRevisionHashLabelKey]; ok {
		return hash, nil
	}
	return "", fmt.Errorf("controller-revision-hash label not found on pod %s", pod.Name)
}
