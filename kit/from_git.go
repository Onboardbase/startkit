package kit

import (
	"fmt"

	"github.com/Onboardbase/obbkitv2/utils"
)

func InitFromGit(input InitFromGitInput) {
	projectName := collectProjectName()
	cloneGitRepo(CloneGitRepoInput{
		GitRepoURL:  input.GitRepoURL,
		ProjectName: projectName,
	})

	SetupOnboardbase(OnboardbaseSetupInput{
		StartCommand:      `echo "Done🙌🏾"`,
		ProjectFolderName: projectName,
	})
}

func cloneGitRepo(input CloneGitRepoInput) (){
	cloneCommand := fmt.Sprintf("git clone %s %s", input.GitRepoURL, input.ProjectName)

	utils.RunShellCommand(utils.RunShellCommandInput{
		ShellToUse:       "bash",
		Command:          cloneCommand,
	})
	return
}

type InitFromGitInput struct {
	GitRepoURL string
}

type CloneGitRepoInput struct {
	GitRepoURL  string
	ProjectName string
}
