package project

import (
	"context"
	"fmt"

	redhatcopv1alpha1 "github.com/redhat-cop/project-initialize-operator/project-initialize/pkg/apis/redhatcop/v1alpha1"
	"github.com/redhat-cop/quay-operator/pkg/controller/quayecosystem/logging"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func AddQuotaToProject(client client.Client, quota *corev1.ResourceQuota) error {
	err := client.Create(context.TODO(), quota)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Failed to create quota %s", quota.ObjectMeta.Name))
		return err
	}

	logging.Log.Info(fmt.Sprintf("Created quota %s for namespace %s", quota.ObjectMeta.Name, quota.ObjectMeta.Namespace))
	return nil
}

func GetQuotaSizeFromCluster(client client.Client, name string) (error, *redhatcopv1alpha1.ProjectInitializeQuota) {
	quotaSize := &redhatcopv1alpha1.ProjectInitializeQuota{}
	err := client.Get(context.TODO(), types.NamespacedName{Name: name}, quotaSize)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Quota definition not found: %s", name))
		return err, nil
	}

	return nil, quotaSize
}
