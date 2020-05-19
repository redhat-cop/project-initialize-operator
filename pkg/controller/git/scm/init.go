package scm

func GitInit(teamName string) error {
	/* TODO - push an example boilerplate project to a repo gitops-${teamName}*/
	return nil
}

func CreateRepoBitBucket(teamName string) error {
	/* TODO - create a new repo if one does not already exist on bitbucket*/
	/* You will need a provlidged service account on bitbucket that has this access and use the API to create a repo */
	/* Store credentials as secret*/
	return nil
}

func CreateRepoGitHub(teamName string) error {
	/* TODO - create a new repo if one does not already exist on github*/
	/* You will need a provlidged service account on github that has this access and use the API to create a repo */
	/* Store credentials as secret*/
	return nil
}
