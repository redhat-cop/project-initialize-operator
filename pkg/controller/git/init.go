package git

import (
	"context"
	"errors"
	"time"

	redhatcopv1alpha1 "github.com/redhat-cop/project-initialize-operator/project-initialize/pkg/apis/redhatcop/v1alpha1"
	github "github.com/redhat-cop/project-initialize-operator/project-initialize/pkg/controller/git/github"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("git")

func GitInitialize(c client.Client, namespace string, teamName string, env string, git *redhatcopv1alpha1.Git, gitTemplate *redhatcopv1alpha1.GitTemplate) error {
	if git.Provider == redhatcopv1alpha1.GitHub {
		err := createRepoGitHub(c, teamName, namespace, env, git, gitTemplate)
		if err != nil {
			return err
		}
		return nil
	}
	if git.Provider == redhatcopv1alpha1.BitBucket {
		createRepoBitBucket(teamName)
		return nil
	}
	if git.Provider == redhatcopv1alpha1.GitLab {
		createRepoGitLab(teamName)
		return nil
	}
	return errors.New("Invalid GIT Host Type")
}

func createRepoBitBucket(teamName string) error {
	return errors.New("BitBucket not available yet")
}

func createRepoGitHub(c client.Client, teamName string, namespace string, env string, git *redhatcopv1alpha1.Git, gitTemplate *redhatcopv1alpha1.GitTemplate) error {
	tokenSecret := &corev1.Secret{}
	err := c.Get(context.TODO(), types.NamespacedName{Name: git.AccountSecret.Name, Namespace: git.AccountSecret.Namespace}, tokenSecret)
	if err != nil {
		return err
	}
	token := string(tokenSecret.Data["token"])
	gitClient := github.InitializeGitHubClient(token)

	hasGit, err := github.CheckForGitOpsRepository(gitClient, git.Suffix, teamName, git.Owner)
	if err != nil {
		return err
	}
	if !hasGit {
		// Check if there is a template provided or if this should just be a new blank repository
		if gitTemplate != nil {
			github.CreateNewRespositoryWithTemplate(gitClient, git.Suffix, teamName, gitTemplate.Owner, gitTemplate.Repo, git)
		} else {
			github.CreateNewRespository(gitClient, git.Suffix, teamName, git)
		}
	}

	time.Sleep(10 * time.Second)
	err = github.GitAddEnvironment(gitClient, env, git.Owner, github.GetTeamRepoName(teamName, git.Suffix))
	if err != nil {
		return err
	}
	return nil
}

func createRepoGitLab(teamName string) error {
	return errors.New("GitLab not available yet")
}
