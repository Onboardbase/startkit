package kit

import (
	"fmt"

	"github.com/Onboardbase/obbkit/utils"
)

type initiationCommands []string

func (ic initiationCommands) Run(dirToRunIn string) {
	if len(ic) == 0 {
		return
	}
	fmt.Println("\n---------------------------------------------------------------------")
	fmt.Println("|\t\U000023F3 Running steps to configure project...")
	fmt.Println("---------------------------------------------------------------------")
	for _, command := range ic {
		fmt.Println("🔧", command)
		utils.RunShellCommand(utils.RunShellCommandInput{
			ShellToUse:       "bash",
			Command:          command,
			DirectoryToRunIn: dirToRunIn,
		})
	}
}

type initFrameworkProjectInput struct {
	ProjectName                 string
	ShellCommandToCreateProject string
	StartCommand                string
	FrameworkName               string
	FirstStepsCommands          initiationCommands
}

type createFrameworkProjectInput struct {
	ShellCommandToCreateProject string
	ProjectName                 string
	FrameworkName               string
	FirstStepsCommands          initiationCommands
}

func initFrameworkProject(input initFrameworkProjectInput) {
	projectFolderName := createFrameworkProject(createFrameworkProjectInput{
		FrameworkName:               input.FrameworkName,
		ShellCommandToCreateProject: input.ShellCommandToCreateProject,
		ProjectName:                 input.ProjectName,
	})

	setupOnboardbase(onboardbaseSetupInput{
		StartCommand:       input.StartCommand,
		ProjectFolderName:  projectFolderName,
		FirstStepsCommands: input.FirstStepsCommands,
	})
}

func createFrameworkProject(input createFrameworkProjectInput) (projectFolderName string) {
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

type onboardbaseSetupInput struct {
	StartCommand       string
	ProjectFolderName  string
	FirstStepsCommands initiationCommands
}

func setupOnboardbase(input onboardbaseSetupInput) error {
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
