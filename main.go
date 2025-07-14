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
}

func main() {
	fmt.Println("Reading from config.yml")
	b, err := os.ReadFile("config.yml")
	if err != nil {
		log.Println(err)
	}

	workloads := make(map[string][]workloadObject)
	err = yaml.Unmarshal(b, &workloads)
	if err != nil {
		log.Print(err)
	}

	for _, groupObjects := range workloads {
		//fmt.Printf("Workload group: %s\n", groupName)
		for _, obj := range groupObjects {
			if obj.CreateNamespace {
				k8s.CreateNamespace(obj.Namespace)
			}
		}
	}
}
