package kit

import (
	"fmt"

	"github.com/Onboardbase/startkit/utils"
)

type InitFromGitInput struct {
	GitRepoURL string
}

type cloneGitRepoInput struct {
	GitRepoURL  string
	ProjectName string
}

func InitFromGit(input InitFromGitInput) {
	projectName := collectProjectName()
	gitCloneAndSetupOnboardbase(input, projectName)
}

func gitCloneAndSetupOnboardbase(input InitFromGitInput, projectName string) {
	cloneGitRepo(cloneGitRepoInput{
		GitRepoURL:  input.GitRepoURL,
		ProjectName: projectName,
	})
	setupOnboardbase(onboardbaseSetupInput{
		StartCommand:      `echo "Done🙌🏾"`,
		ProjectFolderName: projectName,
	})
}

func cloneGitRepo(input cloneGitRepoInput) {
	cloneCommand := fmt.Sprintf("git clone %s %s", input.GitRepoURL, input.ProjectName)
	utils.RunShellCommand(utils.RunShellCommandInput{
		ShellToUse: "bash",
		Command:    cloneCommand,
	})
}
