package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ramalabeysekera/k8s-deploy-scaffold/helpers"
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
	verbose := flag.Bool("verbose", false, "Enable verbose output")
	// Add kubeconfig flag
	kubeconfig := flag.String("kubeconfig", "", "(optional) absolute path to the kubeconfig file")
	flag.Parse()
	helpers.Verbose = *verbose

	// If kubeconfig is not set, use default path
	if *kubeconfig == "" {
		home, _ := os.UserHomeDir()
		*kubeconfig = fmt.Sprintf("%s/.kube/config", home)
	}
	k8s.InitConfig(*kubeconfig)

	fmt.Println("=======================")
	fmt.Println("🚀 K8s Deploy Scaffold")
	fmt.Println("=======================")
	fmt.Println("📖 Reading configuration from: config.yml")
	fmt.Println("=========================================")
	b, err := os.ReadFile("config.yml")
	if err != nil {
		fmt.Printf("❌ Error: failed to read config.yml: %v\n", err)
		return
	}

	workloads := make(map[string][]workloadObject)
	err = yaml.Unmarshal(b, &workloads)
	if err != nil {
		fmt.Printf("❌ Error: failed to unmarshal config: %v\n", err)
		return
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
					fmt.Printf("❌ Error: %v\n", err)
				}
			}
			fmt.Println("==============================")
		}
	}
}
