package main

import "fmt"

func main() {

	switch "asdf" {
	case "kubernetes":
		fmt.Println("Case 1. kuberenetes with lower-case \"k\".")
	case "Kubernetes":
		fmt.Println("Case 2. kuberenetes with upper-case \"K\".")
	case "K8s":
		fmt.Println("Case 3. Kubernetes short name \"kates\".")
	case "Istio":
		fmt.Println("Case 4. Service mesh is the future.")
	default:
		fmt.Println("Hit the default")
	}
}
