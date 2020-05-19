package project

import (
	"context"
	"fmt"

	redhatcopv1alpha1 "github.com/redhat-cop/project-initialize-operator/pkg/apis/redhatcop/v1alpha1"
	"github.com/redhat-cop/project-initialize-operator/pkg/controller/logging"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func UpdateNamespaceAnnotations(client client.Client, name string, namespaceDetails *redhatcopv1alpha1.NamespaceDetails) error {
	ns := &corev1.Namespace{}

	err := client.Get(context.TODO(), types.NamespacedName{Name: name}, ns)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Failed to find namespace %s", name))
		return err
	}

	// OCP adds annotations at creation, we need to copy over new values
	if namespaceDetails.Annotations != nil {
		for key, value := range namespaceDetails.Annotations {
			ns.Annotations[key] = value
		}
	}

	if namespaceDetails.Annotations != nil {
		ns.Labels = namespaceDetails.Labels
	}

	err = client.Update(context.TODO(), ns)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Failed to update namespace %s", name))
		return err
	}

	logging.Log.Info(fmt.Sprintf("Updated namespace %s", name))
	return nil
}
