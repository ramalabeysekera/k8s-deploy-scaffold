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
	fmt.Printf("🔧 Creating Namespace: %s\n", namespace)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("❌ Error: could not create Kubernetes client: %v\n", err)
		panic(err.Error())
	}

	obj := metav1.ObjectMeta{
		Name: namespace,
	}

	namespaceObj, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{ObjectMeta: obj}, metav1.CreateOptions{})

	if errors.IsAlreadyExists(err) {
		fmt.Printf("⚠️ Namespace '%s' already exists\n", namespace)
		return nil
	} else if err != nil {
		fmt.Printf("❌ Error: failed to create namespace '%s': %v\n", namespace, err)
		return err
	}
	fmt.Printf("✅ Created Namespace '%s' at %v\n", namespace, namespaceObj.CreationTimestamp)
	return nil
}
