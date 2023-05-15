package kit

import (
	"fmt"
	"os"
	"path"

	"github.com/Onboardbase/obbkitv2/utils"
)

type OnboardbaseSetupInput struct {
	StartCommand string
	ProjectFolderName  string
}

func SetupOnboardbase(input OnboardbaseSetupInput) error {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	projectDir := path.Join(currentDir, input.ProjectFolderName)
	cmd := fmt.Sprintf("onboardbase auth -c \"%s\"", input.StartCommand)
	err = utils.RunShellCommand(utils.RunShellCommandInput{
		ShellToUse: "bash",
		Command:    cmd,
		DirectoryToRunIn:  projectDir,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}