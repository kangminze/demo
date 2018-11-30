package main

import "demo/src/kubernetes"

func main() {
	client, _ := kubernetes.NewClient()
	client.GetVirtualServices("default", "details")
}
