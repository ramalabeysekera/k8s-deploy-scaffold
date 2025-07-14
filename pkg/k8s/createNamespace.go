package k8s

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateNamespace(namespace string) error {
	fmt.Printf("Creating namespace: %s\n", namespace)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	obj := metav1.ObjectMeta{
		Name: namespace,
	}

	namespaceObj, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{ObjectMeta: obj}, metav1.CreateOptions{})

	if errors.IsAlreadyExists(err) {
		fmt.Printf("Namespace %s already exist\n", namespace)
		return nil
	} else if err != nil {
		return err
	}
	fmt.Printf("Created namespace %s at %v\n", namespace, namespaceObj.CreationTimestamp)
	return nil
}
