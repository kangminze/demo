package kubernetes

import (
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (client *IstioClient) GetServices(namespace string, selectorLabels map[string]string) ([]v1.Service, error) {
	allServiceList, err := client.k8s.CoreV1().Services(namespace).List(meta_v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var services []v1.Service = allServiceList.Items

	return services, nil
}

func (client *IstioClient) GetPods(namespaces string, selectorLabels map[string]string) ([]v1.Pod, error) {
	allPod, err := client.k8s.CoreV1().Pods(namespaces).List(meta_v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var pods []v1.Pod = allPod.Items
	return pods, nil
}
