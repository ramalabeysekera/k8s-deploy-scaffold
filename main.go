package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ramalabeysekera/k8s-deploy-scaffold/pkg/k8s"
	"gopkg.in/yaml.v3"
)

type workloadObject struct {
	Namespace            string `yaml:"namespace"`
	CreateNamespace      bool   `yaml:"create_namespace"`
	CreateServiceAccount bool   `yaml:"create_service_account"`
	ServiceAccountName   string `yaml:"service_account_name"`
	EnableIRSA           bool   `yaml:"enable_irsa"`
	OIDCProvider         string `yaml:"oidc_provider"`
}

func main() {
	fmt.Println("=======================")
	fmt.Println("🚀 K8s Deploy Scaffold")
	fmt.Println("=======================")
	fmt.Println("📖 Reading configuration from: config.yml")
	fmt.Println("=========================================")
	b, err := os.ReadFile("config.yml")
	if err != nil {
		log.Println(err)
	}

	workloads := make(map[string][]workloadObject)
	err = yaml.Unmarshal(b, &workloads)
	if err != nil {
		log.Print(err)
	}

	for groupName, groupObjects := range workloads {
		for _, obj := range groupObjects {

			fmt.Printf("Workload: %s\n", groupName)
			if obj.CreateNamespace {
				k8s.CreateNamespace(obj.Namespace)
			}

			if obj.CreateServiceAccount {
				err := k8s.CreateServiceAccount(obj.ServiceAccountName, obj.Namespace, obj.EnableIRSA, obj.OIDCProvider)

				if err != nil {
					fmt.Println(err)
				}
			}
			fmt.Println("==============================")
		}
	}
}
