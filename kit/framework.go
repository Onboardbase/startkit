package kit

import (
	"fmt"
	"os"
	"path"

	"github.com/Onboardbase/obbkitv2/utils"
)

type FrameworkHandler interface {
	CreateProject(input CreateFrameworkProjectInput) (projectFolderName string) // returns the name of the folder the project was initialized in
}

type CreateFrameworkProjectInput struct {
	ShellCommandToCreateProject string
	ProjectName  string
}

type SetupOnboardbaseInput struct {
	StartCommand string
	ProjectFolderName  string
}
func (h *nestJSHandler) SetupOnboardbase(input SetupOnboardbaseInput) error {
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