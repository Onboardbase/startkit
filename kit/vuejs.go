package kit

import (
	"fmt"
)

func initVueProject(projectName string) {
	createProjectCmd := fmt.Sprintf("npm create vue@3 %s", projectName)

	InitFrameworkProject(InitFrameworkProjectInput{
		ProjectName:                 projectName,
		ShellCommandToCreateProject: createProjectCmd,
		StartCommand:                "yarn start",
		FrameworkName:               "Vue",
	})
}
