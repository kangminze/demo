package service

import (
	"demo/src/kubernetes"
	"demo/src/model"
	"github.com/sevenNt/wzap"
	"k8s.io/api/core/v1"
	"sync"
)

type SvcService struct {
	k8s kubernetes.IstioClientInterface
}

func NewSvcService() *SvcService {
	istioClient, _ := kubernetes.NewClient()
	return &SvcService{
		k8s: istioClient,
	}
}

//func NewSvcService(k8s kubernetes.IstioClientInterface) *SvcService{
//	return &SvcService{
//		k8s:k8s,
//	}
//}

func (svc *SvcService) GetServiceList(namespace string) (*model.ServiceList, error) {

	var (
		allService  []v1.Service
		allPods     []v1.Pod
		wg          sync.WaitGroup
		services    []model.ServiceOverview
		serviceList *model.ServiceList
	)
	wg.Add(2)
	errChan := make(chan error, 2)
	go func() {
		defer wg.Done()
		var err error
		allService, err = svc.k8s.GetServices(namespace, nil)
		if err != nil {
			wzap.Err(err)
			errChan <- err
		}

	}()

	go func() {
		defer wg.Done()
		var err error
		allPods, err = svc.k8s.GetPods(namespace, nil)
		if err != nil {
			wzap.Err(err)
			errChan <- err
		}
	}()
	wg.Wait()
	if len(errChan) > 0 {
		err := <-errChan
		return nil, err
	}

	//为service过滤出对应的pod
	services = make([]model.ServiceOverview, len(allService))
	for i, svc := range allService {

		svcPods, _ := filterSvcPods(svc, allPods)
		// 判断pod annotation是否包含istio相关key
		hasIstio := HasIstio(svcPods)
		services[i] = model.ServiceOverview{
			Name:         svc.Name,
			IstioSidecar: hasIstio,
		}
	}
	serviceList = &model.ServiceList{
		Namespace: namespace,
		Services:  services,
	}
	wzap.Debug("successs")
	return serviceList, nil

}

// 判断pod是包含istio组件
func HasIstio(pods []v1.Pod) bool {
	for _, pod := range pods {
		if _, ok := pod.Annotations["sidecar.istio.io/status"]; ok {
			return true
		}
	}
	return false
}

func buildServiceList() {

}
