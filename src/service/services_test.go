package service

import (
	"demo/src/kubernetes"
	"testing"
)

func TestServiceList(t *testing.T) {
	istioClient, _ := kubernetes.NewClient()
	svcService := SvcService{istioClient}
	svcService.GetServiceList("default")
}
