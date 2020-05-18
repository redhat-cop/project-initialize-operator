package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v31/github"
	redhatcopv1alpha1 "github.com/redhat-cop/project-initialize-operator/project-initialize/pkg/apis/redhatcop/v1alpha1"
)

func CheckForGitOpsRepository(client *github.Client, suffix string, teamName string, owner string) (bool, error) {
	// Check if repo exists
	repoName := GetTeamRepoName(teamName, suffix)

	repos, _, err := client.Repositories.List(context.Background(), owner, nil)
	if err != nil {
		return false, err
	}
	// Search repos to see if the teams repo exists
	for _, repo := range repos {
		if repo.GetName() == repoName {
			return true, nil
		}
	}
	return false, nil
}
func CreateNewRespository(client *github.Client, suffix string, teamName string, gitDetails *redhatcopv1alpha1.Git) error {
	repoName := GetTeamRepoName(teamName, suffix)
	newRepo := getRepoRequest(repoName, gitDetails.Owner, gitDetails.Private, gitDetails.Desc)
	repo, _, err := client.Repositories.Create(context.TODO(), "", newRepo)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("Successfully created new repo: %v\n", repo.GetName()))
	return nil
}

func CreateNewRespositoryWithTemplate(client *github.Client, suffix string, teamName string, templateOwner string, templateRepo string, gitDetails *redhatcopv1alpha1.Git) error {

	repoName := GetTeamRepoName(teamName, suffix)
	newRepo := getTemplateRequest(repoName, gitDetails.Owner, gitDetails.Private, gitDetails.Desc)
	repo, _, err := client.Repositories.CreateFromTemplate(context.Background(), "", templateRepo, newRepo)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("Successfully created new repo: %v\n", repo.GetName()))
	return nil
}

func getTemplateRequest(repoName string, owner string, private bool, description string) *github.TemplateRepoRequest {
	return &github.TemplateRepoRequest{
		Name:        &repoName,
		Owner:       &owner,
		Private:     &private,
		Description: &description,
	}
}

func getRepoRequest(repoName string, owner string, private bool, description string) *github.Repository {
	autoInit := true

	return &github.Repository{
		Name:        &repoName,
		Private:     &private,
		Description: &description,
		AutoInit:    &autoInit,
	}
}

func GetTeamRepoName(teamName string, suffix string) string {
	return fmt.Sprintf("%s-%s", teamName, suffix)
}
