package service

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

//过滤出改svc下所有pods
func filterSvcPods(service v1.Service, allPods []v1.Pod) ([]v1.Pod, error) {
	var pods []v1.Pod
	svcSelector := labels.Set(service.Spec.Selector).AsSelector()
	for _, pod := range allPods {

		if svcSelector.Matches(labels.Set(pod.ObjectMeta.Labels)) {
			pods = append(pods, pod)
		}
	}
	return pods, nil
}
