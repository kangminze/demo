package kubernetes

import (
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

type IstioClientInterface interface {
	GetVirtualServices(namespace string, serviceName string) ([]IstioObject, error)
	GetServices(namespace string, selectorLabels map[string]string) ([]v1.Service, error)
	GetPods(namespaces string, selectorLabels map[string]string) ([]v1.Pod, error)
}

type IstioClient struct {
	IstioClientInterface
	k8s                *kubernetes.Clientset
	istioConfigApi     *rest.RESTClient
	istioNetworkingApi *rest.RESTClient
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	return os.Getenv("USERPROFILE")
}

func ConfigClient() (*rest.Config, error) {
	var kubeconfig string
	home := homeDir()
	if len(home) <= 0 {
		return nil, errors.New("home is nil")
	}
	kubeconfig = filepath.Join(home, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func NewClient() (*IstioClient, error) {
	config, err := ConfigClient()
	if err != nil {
		return nil, err
	}
	client := IstioClient{}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	client.k8s = clientSet

	//Istio CRD扩展自kubernetes apis
	types := runtime.NewScheme()
	schemeBuilder := runtime.NewSchemeBuilder(
		func(scheme *runtime.Scheme) error {
			for _, nt := range networkingTypes {
				scheme.AddKnownTypeWithName(networkingGroupVersion.WithKind(nt.objectKind), &GenericIstioObject{})
				scheme.AddKnownTypeWithName(networkingGroupVersion.WithKind(nt.collectionKind), &GenericIstioObjectList{})
			}

			for _, cf := range configTypes {
				scheme.AddKnownTypeWithName(configGroupVersion.WithKind(cf.objectKind), &GenericIstioObject{})
				scheme.AddKnownTypeWithName(configGroupVersion.WithKind(cf.collectionKind), &GenericIstioObjectList{})
			}

			for _, ad := range adapterTypes {
				scheme.AddKnownTypeWithName(configGroupVersion.WithKind(ad.objectKind), &GenericIstioObject{})
				scheme.AddKnownTypeWithName(configGroupVersion.WithKind(ad.collectionKind), &GenericIstioObjectList{})
			}

			for _, tp := range templateTypes {
				scheme.AddKnownTypeWithName(configGroupVersion.WithKind(tp.objectKind), &GenericIstioObject{})
				scheme.AddKnownTypeWithName(configGroupVersion.WithKind(tp.collectionKind), &GenericIstioObjectList{})
			}
			meta_v1.AddToGroupVersion(scheme, configGroupVersion)
			meta_v1.AddToGroupVersion(scheme, networkingGroupVersion)
			return nil
		})
	err = schemeBuilder.AddToScheme(types)
	if err != nil {
		return nil, err
	}

	//istio config
	istioConfig := rest.Config{
		Host:    config.Host,
		APIPath: "/apis",
		ContentConfig: rest.ContentConfig{
			GroupVersion:         &configGroupVersion,
			NegotiatedSerializer: serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(types)},
			ContentType:          runtime.ContentTypeJSON,
		},
		Burst:           config.Burst,
		TLSClientConfig: config.TLSClientConfig,
		BearerToken:     config.BearerToken,
		QPS:             config.QPS,
	}
	istioConfigApi, err := rest.RESTClientFor(&istioConfig)
	if err != nil {
		return nil, err
	}
	client.istioConfigApi = istioConfigApi
	istioNetworkingConfig := rest.Config{
		Host:    config.Host,
		APIPath: "/apis",
		ContentConfig: rest.ContentConfig{
			GroupVersion:         &networkingGroupVersion,
			NegotiatedSerializer: serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(types)},
			ContentType:          runtime.ContentTypeJSON,
		},
		Burst:           config.Burst,
		TLSClientConfig: config.TLSClientConfig,
		BearerToken:     config.BearerToken,
		QPS:             config.QPS,
	}

	istioNetworkingApi, err := rest.RESTClientFor(&istioNetworkingConfig)

	if err != nil {
		return nil, err
	}
	client.istioNetworkingApi = istioNetworkingApi
	return &client, nil
}
