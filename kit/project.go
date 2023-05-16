package kit

import (
	"fmt"

	"github.com/Onboardbase/obbkitv2/utils"
)

const (
	Nestjs  = "NestJS"
	Nextjs  = "NextJS"
	Reactjs = "React.js"
	Vuejs   = "VueJs"
)

func Init() {
	projectName, projectType := askForProjectNameAndType()

	switch projectType {
	case Nestjs:
		initNestJsProject(projectName)
	case Nextjs:
		initNextProject(projectName)
	case Reactjs:
		initReactProject(projectName)
	case Vuejs:
		initVueProject(projectName)
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
	items := []string{Nestjs, Nextjs, Reactjs, Vuejs}
	projectType := utils.GetSelectInputFromPrompt(projectTypePromptContent, items)
	return projectType
}

