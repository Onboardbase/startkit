package kit

import (
	"fmt"
)

func initReactProject(projectName string) {
	createProjectCmd := fmt.Sprintf("npx create-react-app %s", projectName)
	initFrameworkProject(initFrameworkProjectInput{
		ProjectName:                 projectName,
		ShellCommandToCreateProject: createProjectCmd,
		StartCommand:                "npm run start",
		FrameworkName:               "React",
	})
}
