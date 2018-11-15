package client

import (
	"github.com/sevenNt/wzap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func Init() {
	var kubeconfig string

	if home := homeDir(); home != "" {

		kubeconfig = filepath.Join(home, ".kube", "config")


	} else {
		panic("找不到配置文件")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset,err := kubernetes.NewForConfig(config)
	if err != nil {
		wzap.Err(err)
		panic(err)
	}
	namespaces,err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	wzap.Info(namespaces.String())
}


func homeDir() string  {

	if home := os.Getenv("HOME"); home != "" {
		return home
	}

	return os.Getenv("USERPROFILE")

}