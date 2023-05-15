package kit

import (
	"fmt"

	"github.com/Onboardbase/obbkitv2/utils"
)

type FrameworkHandler interface {
	CreateProject(input CreateFrameworkProjectInput) (projectFolderName string) // returns the name of the folder the project was initialized in
}

type CreateFrameworkProjectInput struct {
	ShellCommandToCreateProject string
	ProjectName                 string
	FrameworkName               string
}

type InitFrameworkProjectInput struct {
	ProjectName                 string
	ShellCommandToCreateProject string
	StartCommand                string
	FrameworkName               string
}

func InitFrameworkProject(input InitFrameworkProjectInput) {
	projectFolderName := CreateFrameworkProject(CreateFrameworkProjectInput{
		FrameworkName:               input.FrameworkName,
		ShellCommandToCreateProject: input.ShellCommandToCreateProject,
		ProjectName:                 input.ProjectName,
	})

	SetupOnboardbase(OnboardbaseSetupInput{
		StartCommand:      input.StartCommand,
		ProjectFolderName: projectFolderName,
	})
}

func CreateFrameworkProject(input CreateFrameworkProjectInput) (projectFolderName string) {
	message := fmt.Sprintf("|\t\U000023F3 Creating %s project...", input.FrameworkName)
	fmt.Println("\n---------------------------------------------------------------------")
	fmt.Println(message)
	fmt.Println("---------------------------------------------------------------------")

	err := utils.RunShellCommand(utils.RunShellCommandInput{
		ShellToUse: "bash",
		Command:    input.ShellCommandToCreateProject,
	})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	projectFolderName = input.ProjectName
	return
}
