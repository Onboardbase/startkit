package kit

import (
	"fmt"

	"github.com/Onboardbase/obbkitv2/utils"
)

type nestJSHandler struct {
}

func (h *nestJSHandler) CreateProject(input CreateFrameworkProjectInput) (projectFolderName string) {
	fmt.Println("\t \U000023F3 Creating NestJS project...")
	cmd := fmt.Sprintf("npx @nestjs/cli new %s", input.ProjectName)
	err := utils.RunShellCommand(utils.RunShellCommandInput{
		ShellToUse: "bash",
		Command:    cmd,
	})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	projectFolderName = utils.NormalizeToKebabOrSnakeCase(input.ProjectName)
	return
}

func initNestJsProject(projectName string) {
	handler := &nestJSHandler{}

	createProjectCmd := fmt.Sprintf("npx @nestjs/cli new %s", projectName)
	projectFolderName := handler.CreateProject(CreateFrameworkProjectInput{
		ProjectName:                 projectName,
		ShellCommandToCreateProject: createProjectCmd,
	})

	startProjectCmd := "yarn start"
	handler.SetupOnboardbase(SetupOnboardbaseInput{
		StartCommand:      startProjectCmd,
		ProjectFolderName: projectFolderName,
	})
}
