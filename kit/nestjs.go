package kit

import (
	"fmt"
)

func initNestJsProject(projectName string) {
	createProjectCmd := fmt.Sprintf("npx @nestjs/cli new %s", projectName)
	initFrameworkProject(initFrameworkProjectInput{
		ProjectName:                 projectName,
		ShellCommandToCreateProject: createProjectCmd,
		StartCommand:                "npm run start",
		FrameworkName:               "Nest",
	})
}
