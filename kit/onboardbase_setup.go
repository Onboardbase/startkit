package kit

import (
	"fmt"

	"github.com/Onboardbase/obbkitv2/utils"
)

type OnboardbaseSetupInput struct {
	StartCommand       string
	ProjectFolderName  string
	FirstStepsCommands initiationCommands
}

func SetupOnboardbase(input OnboardbaseSetupInput) error {
	projectDir := utils.GetProjectPath(input.ProjectFolderName)
	input.FirstStepsCommands.Run(projectDir)
	cmd := fmt.Sprintf("onboardbase auth -c \"%s\" --overwrite", input.StartCommand)
	err := utils.RunShellCommand(utils.RunShellCommandInput{
		ShellToUse:       "bash",
		Command:          cmd,
		DirectoryToRunIn: projectDir,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
