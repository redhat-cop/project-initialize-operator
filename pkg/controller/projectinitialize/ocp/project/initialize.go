package project

import (
	"fmt"

	projectv1 "github.com/openshift/api/project/v1"
	projectset "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	redhatcopv1alpha1 "github.com/redhat-cop/project-initialize-operator/pkg/apis/redhatcop/v1alpha1"
	"github.com/redhat-cop/project-initialize-operator/pkg/controller/logging"
)

func InitializeProjectOCP(client *projectset.ProjectV1Client, projectRequest *projectv1.ProjectRequest) (*projectv1.Project, error) {
	project, err := client.ProjectRequests().Create(projectRequest)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Failed to create project %s", projectRequest.Name))
		return nil, err
	}

	return project, nil
}

func GetProjectName(projectInitialize *redhatcopv1alpha1.ProjectInitialize) string {

	if projectInitialize.Spec.NamespaceDetails != nil && projectInitialize.Spec.NamespaceDetails.Name != "" {
		return projectInitialize.Spec.NamespaceDetails.Name
	} else {
		return fmt.Sprintf("%s-%s", projectInitialize.Spec.Team, projectInitialize.Spec.Env)
	}
}
