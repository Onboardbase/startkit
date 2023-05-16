package kit

import (
	"fmt"
)

func initNestJsProject(projectName string) {
	createProjectCmd := fmt.Sprintf("npx @nestjs/cli new %s", projectName)

	InitFrameworkProject(InitFrameworkProjectInput{
		ProjectName:                 projectName,
		ShellCommandToCreateProject: createProjectCmd,
		StartCommand:                "npm run start",
		FrameworkName:               "Nest",
	})
}