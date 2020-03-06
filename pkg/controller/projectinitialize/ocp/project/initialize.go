package project

import (
	"fmt"

	projectv1 "github.com/openshift/api/project/v1"
	projectset "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	"github.com/redhat-cop/quay-operator/pkg/controller/quayecosystem/logging"
)

func InitializeProjectOCP(client *projectset.ProjectV1Client, projectRequest *projectv1.ProjectRequest) (*projectv1.Project, error) {
	project, err := client.ProjectRequests().Create(projectRequest)
	if err != nil {
		logging.Log.Error(err, fmt.Sprintf("Failed to create project %s", projectRequest.Name))
		return nil, err
	}

	return project, nil
}

func GetProjectName(team string, env string) string {
	return fmt.Sprintf("%s-%s", team, env)
}
