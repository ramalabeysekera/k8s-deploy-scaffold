package k8s

import (
	"fmt"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var config *rest.Config
var err error

// InitConfig initializes the k8s config with the provided kubeconfig path
func InitConfig(kubeconfigPath string) {
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		fmt.Printf("❌ Error: failed to build kubeconfig: %v\n", err)
		panic(err.Error())
	}
}
