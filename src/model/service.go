package model

type ServiceOverview struct {
	Name         string `json:"name"`
	IstioSidecar bool   `json:"istioSidecar"`
}

type ServiceList struct {
	Namespace string
	Services  []ServiceOverview
}
