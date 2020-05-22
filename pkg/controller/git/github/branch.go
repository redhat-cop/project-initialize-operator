package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/v31/github"
)

func GitAddEnvironment(client *github.Client, env string, owner string, repo string) error {
	found, err := getBranch(client, env, owner, repo)
	if err != nil {
		return err
	}

	if found == nil {
		master, err := getBranch(client, "master", owner, repo)
		if err != nil {
			return err
		}
		if master == nil || master.Commit.SHA == nil {
			return errors.New(fmt.Sprintf("SHA1 for master branch on respository %s not found", repo))
		}
		emptyString := ""
		ref := fmt.Sprintf("refs/heads/%s", env)
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

func getBranch(client *github.Client, branch string, owner string, repo string) (*github.Branch, error) {
	ret, res, _ := client.Repositories.GetBranch(context.TODO(), owner, repo, branch)

	if res.StatusCode == 200 {

		return ret, nil
	}

	return nil, nil
}
