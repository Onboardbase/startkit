package kit

func initFlaskProject(projectName string) {
	gitCloneAndSetupOnboardbase(InitFromGitInput{
		GitRepoURL: "https://github.com/Onboardbase/Flask-Starterkit.git",
	}, projectName)
}