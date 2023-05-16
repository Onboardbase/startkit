package kit

import (
	"fmt"
)

func initNextProject(projectName string) {
	createProjectCmd := fmt.Sprintf("npx create-next-app@latest %s", projectName)
	InitFrameworkProject(InitFrameworkProjectInput{
		ProjectName:                 projectName,
		ShellCommandToCreateProject: createProjectCmd,
		StartCommand:                "npm run start",
		FrameworkName:               "Next",
	})
}
