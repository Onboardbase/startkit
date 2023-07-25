package kit

import (
	"fmt"

	"github.com/Onboardbase/startkit/project_types"
	"github.com/Onboardbase/startkit/utils"
)

func Init() {
	projectName, projectType := askForProjectNameAndType()

	switch projectType {
	case project_types.Nestjs:
		initNestJsProject(projectName)
	case project_types.Nextjs:
		initNextProject(projectName)
	case project_types.Reactjs:
		initReactProject(projectName)
	case project_types.Vuejs:
		initVueProject(projectName)
	case project_types.Flask:
		initFlaskProject(projectName)
	default:
		fmt.Println("Project type not supported yet")
	}
}

func askForProjectNameAndType() (string, string) {
	projectName := collectProjectName()
	projectType := collectProjectType()
	return projectName, projectType
}

func collectProjectName() string {
	projectNamePromptContent := utils.NewPromptContent(utils.PromptContentInput{
		Label:    "\U000023F3 Please provide a project name (no spaces allowed):",
		ErrorMsg: "\U0000274C You must provide a project name",
	})
	projectName := utils.GetInputFromPrompt(projectNamePromptContent)
	projectName = utils.NormalizeToKebabOrSnakeCase(projectName)
	return projectName
}

func collectProjectType() string {
	projectTypePromptContent := utils.NewPromptContent(utils.PromptContentInput{
		Label:    "\U000023F3 Please select a project type.",
		ErrorMsg: "\U0000274C You must select a project type",
	})
	items := []string {
		project_types.Nestjs,
		project_types.Nextjs,
		project_types.Reactjs,
		project_types.Vuejs,
		project_types.Flask,
	}
	projectType := utils.GetSelectInputFromPrompt(projectTypePromptContent, items)
	return projectType
}
