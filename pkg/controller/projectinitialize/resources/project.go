package resources

import (
	projectv1 "github.com/openshift/api/project/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetProjectRequest(name string, displayName string, desc string) *projectv1.ProjectRequest {
	projectrequest := &projectv1.ProjectRequest{
		TypeMeta: metav1.TypeMeta{
			APIVersion: projectv1.SchemeGroupVersion.String(),
			Kind:       "ProjectRequest",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: " ",
		},
		DisplayName: displayName,
		Description: desc,
	}

	return projectrequest
}
