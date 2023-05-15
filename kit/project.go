package kit

import (
	"github.com/Onboardbase/obbkitv2/utils"
)

const (
	Nestjs = "NestJS"
	Nextjs = "NextJS"
	Reactjs = "React.js"
	Vuejs = "VueJs"
)

func CreateProject() {
	projectName, projectType := askForProjectNameAndType()

	if projectType == Nestjs {
		initNestJsProject(projectName)
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
	return projectName
}

func collectProjectType() string {
	projectTypePromptContent := utils.NewPromptContent(utils.PromptContentInput{
		Label:    "\U000023F3 Please select a project type.",
		ErrorMsg: "\U0000274C You must select a project type",
	})
	items := []string{"NestJS", "NextJS", "React.js", "VueJs"}
	projectType := utils.GetSelectInputFromPrompt(projectTypePromptContent, items)
	return projectType
}
