package kubernetes

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/arjit95/kubepf/pkg/cmdutil"
)

// ListNamespaces returns the names of all the pods in a namespace
func ListNamespaces(f *cmdutil.Factory) ([]string, error) {
	namespaces, err := f.Client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	names := make([]string, len(namespaces.Items))
	for _, namespace := range namespaces.Items {
		names = append(names, namespace.Name)
	}

	return names, nil
}
