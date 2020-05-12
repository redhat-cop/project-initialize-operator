package github 

import (
	"errors"
	"fmt"
	git "github.com/google/go-github/github"
	redhatcopv1alpha1 "github.com/redhat-cop/project-initialize-operator/project-initialize/pkg/apis/redhatcop/v1alpha1"
)

func CheckRepository(teamName string) bool {
	// Check if repo exists
	client := initializeGitHubClient(token)
	repos, _, err := client.Repositories.List(ctx, "", nil)

	// loop through repos
}

func CreateNewRespository(teamName string, token string, gitSource *redhatcopv1alpha1.GitSource) error {
	client := initializeGitHubClient(token)

	repoName := fmt.Sprintf("%s-gitops", team)
	r := &git.Repository{Name: repoName, Private: gitSource.Private, Description: gitsource.Desc}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
	return nil
}

func CloneNewRespository(teamName string, ) error {
	// Clone repo from existing source
	return nil
}

func initializeGitHubClient(token string) *git.client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}