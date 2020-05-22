package project

import (
	"context"
	"fmt"

	redhatcopv1alpha1 "github.com/redhat-cop/project-initialize-operator/pkg/apis/redhatcop/v1alpha1"
	"github.com/redhat-cop/project-initialize-operator/pkg/controller/logging"
	project "github.com/redhat-cop/project-initialize-operator/pkg/controller/projectinitialize/resources"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateQuota(client client.Client, instance *redhatcopv1alpha1.ProjectInitialize) error {
	err, quotaSize := getQuotaSizeFromCluster(client, instance.Spec.QuotaSize)
	if err != nil {
		return err
	}
	quota := project.GetQuotaResource(GetQuotaName(instance.Spec.Team), GetProjectName(instance.Spec.Team, instance.Spec.Env), *quotaSize)
	err = addQuotaToProject(client, quota)
	if err != nil {
		return err
	}
	err = updateQuotaStatus(client, instance)
	if err != nil {
		return err
	}

	return nil
}

func addQuotaToProject(client client.Client, quota *corev1.ResourceQuota) error {
	err := client.Create(context.TODO(), quota)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Failed to create quota %s", quota.ObjectMeta.Name))
		return err
	}

	logging.Log.Info(fmt.Sprintf("Created quota %s for namespace %s", quota.ObjectMeta.Name, quota.ObjectMeta.Namespace))
	return nil
}

func updateQuotaStatus(client client.Client, instance *redhatcopv1alpha1.ProjectInitialize) error {
	instance.Status.CurrentQuota = instance.Spec.QuotaSize
	err := client.Status().Update(context.TODO(), instance)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Unable to update status on ProjectInitialize %s", instance.ObjectMeta.Name))
		return err
	}

	return nil
}
func DeleteQuotaInNamespace(client client.Client, instance *redhatcopv1alpha1.ProjectInitialize) error {
	name := GetQuotaName(instance.Spec.Team)
	resourceQuota := &corev1.ResourceQuota{}
	err := client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: GetProjectName(instance.Spec.Team, instance.Spec.Env)}, resourceQuota)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Unable to find ResourceQuota %s for deletion", name))
		return err
	}
	err = client.Delete(context.TODO(), resourceQuota)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Unable to delete ResourceQuota %s", name))
		return err
	}

	return nil
}

func getQuotaSizeFromCluster(client client.Client, name string) (error, *corev1.ResourceQuotaSpec) {
	quotaSize := &redhatcopv1alpha1.ProjectInitializeQuota{}
	err := client.Get(context.TODO(), types.NamespacedName{Name: name}, quotaSize)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Quota definition not found: %s", name))
		return err, nil
	}

	return nil, &quotaSize.Spec.ResourceQuotaSpec
}

func GetQuotaName(team string) string {
	return fmt.Sprintf("%s-quota", team)
}
