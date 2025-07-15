package k8s

import (
	"context"
	"fmt"

	"github.com/ramalabeysekera/k8s-deploy-scaffold/helpers"
	"github.com/ramalabeysekera/k8s-deploy-scaffold/pkg/aws"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateServiceAccount(serviceAccount string, namespaceName string, enableIRSA bool, oidcProvider string) error {
	fmt.Printf("🔧 Creating Service Account: %s\n", serviceAccount)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var metaObj metav1.ObjectMeta

	if enableIRSA {
		fmt.Println("🔧 Creating IAM Role")
		awsConfig, accountId := helpers.LoadAwsConfig()
		roleArn, err := aws.CreateIamRole(awsConfig, serviceAccount, accountId, namespaceName, oidcProvider)

		if err != nil {
			return err
		}

		annotations := map[string]string{
			"eks.amazonaws.com/role-arn": roleArn,
		}

		metaObj = metav1.ObjectMeta{
			Name:        serviceAccount,
			Annotations: annotations,
		}
	} else {
		metaObj = metav1.ObjectMeta{
			Name: serviceAccount,
		}
	}

	serviceAccountObj, err := clientset.CoreV1().ServiceAccounts(namespaceName).Create(context.TODO(), &v1.ServiceAccount{ObjectMeta: metaObj}, metav1.CreateOptions{})

	if errors.IsAlreadyExists(err) {
		fmt.Printf("⚠️  Service Account '%s' already exists in namespace '%s'\n", serviceAccount, namespaceName)
		return nil
	} else if err != nil {
		return err
	}
	fmt.Printf("✅ Created Service Account '%s' in namespace '%s' at %v\n", serviceAccount, namespaceName, serviceAccountObj.CreationTimestamp)
	return nil
}
