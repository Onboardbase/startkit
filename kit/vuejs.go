package kit

import (
	"fmt"
)

func initVueProject(projectName string) {
	createProjectCmd := fmt.Sprintf("npm create vue@3 %s", projectName)
	initFrameworkProject(initFrameworkProjectInput{
		ProjectName:                 projectName,
		ShellCommandToCreateProject: createProjectCmd,
		StartCommand:                "npm run dev",
		FrameworkName:               "Vue",
		FirstStepsCommands: initiationCommands{
			"npm install",
		},
	})
}
