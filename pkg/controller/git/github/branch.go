package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v31/github"
)

func GitAddEnvironment(token string, env string, owner string, repo string) error {
	found, err := getBranch(token, env, owner, repo)
	if err != nil {
		return err
	}

	if found == nil {
		master, err := getBranch(token, "master", owner, repo)
		if err != nil {
			return err
		}
		emptyString := ""
		ref := fmt.Sprintf("refs/heads/%s", env)
		client := InitializeGitHubClient(token)
		gitRef := &github.Reference{
			Ref: &ref,
			URL: &emptyString,
			Object: &github.GitObject{
				SHA:  master.Commit.SHA,
				Type: &emptyString,
				URL:  &emptyString,
			},
		}
		_, _, err = client.Git.CreateRef(context.TODO(), owner, repo, gitRef)
		if err != nil {
			return err
		}

		log.Info(fmt.Sprintf("Successfully created new branch: %s", env))
	}

	return nil
}

func getBranch(token string, branch string, owner string, repo string) (*github.Branch, error) {
	client := InitializeGitHubClient(token)
	ret, res, _ := client.Repositories.GetBranch(context.TODO(), owner, repo, branch)

	if res.StatusCode == 200 {

		return ret, nil
	}

	return nil, nil
}
