package kubernetes

import "errors"

func (client *IstioClient) GetVirtualServices(namespace string, serviceName string) ([]IstioObject, error) {
	result, err := client.istioNetworkingApi.Get().Namespace(namespace).Resource(virtualServices).Do().Get()
	if err != nil {
		return nil, err
	}

	virtualServiceList, ok := result.(*GenericIstioObjectList)
	if !ok {
		return nil, errors.New("失败")
	}

	virtualServices := make([]IstioObject, 0)
	for _, virtualService := range virtualServiceList.GetItems() {
		virtualServices = append(virtualServices, virtualService.DeepCopyIstioObject())
	}

	return virtualServices, nil

}
