package k8s

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateServiceAccount(serviceAccount string, namespaceName string, enableIRSA bool) error {
	fmt.Printf("Creating service account: %s\n", serviceAccount)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var metaObj metav1.ObjectMeta

	annotations := map[string]string{
		"eks.amazonaws.com/role-arn": "aws:iam::XXXXXXXXXXXXX:role/my-role",
	}

	if enableIRSA {
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
		fmt.Printf("Service account %s already exist\n", serviceAccount)
		return nil
	} else if err != nil {
		return err
	}
	fmt.Printf("Created service account %s at %v\n", serviceAccount, serviceAccountObj.CreationTimestamp)
	return nil
}
