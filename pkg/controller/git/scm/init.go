package git

import (
	"errors"
)

func GitInit(gitHost string, teamName string) error {
	if gitHost == "github" {
		CreateRepoGitHub(teamName)
	}
	if gitHost == "bitbucket" {
		CreateRepoBitBucket(teamName)
	}
	if gitHost == "gitlab" {
		CreateRepoGitLab(teamName)
	}
	return errors.New("Invalid GIT Host Type %s", gitHost)
}

func CreateRepoBitBucket(teamName string) error {
	/* TODO - create a new repo if one does not already exist on bitbucket*/
	/* You will need a provlidged service account on bitbucket that has this access and use the API to create a repo */
	/* Store credentials as secret*/
}

func CreateRepoGitHub(teamName string) error {
	/* TODO - create a new repo if one does not already exist on github*/
	/* You will need a provlidged service account on github that has this access and use the API to create a repo */
	/* Store credentials as secret*/
}

func CreateRepoGitLab(teamName string) error {
	/* TODO - create a new repo if one does not already exist on gitlab*/
	/* You will need a provlidged service account on gitlab that has this access and use the API to create a repo */
	/* Store credentials as secret*/
}
