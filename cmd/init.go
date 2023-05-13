/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Onboardbase/obbkitv2/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "I",
	Run: func(cmd *cobra.Command, args []string) {
		createProject()
	},
}


func createProject() {
    projectNamePromptContent := utils.NewPromptContent(utils.PromptContentInput{
		Label: "⏳ Please provide a project name (no spaces allowed):",
		ErrorMsg: "❌ You must provide a project name",
	})

    projectName := utils.GetInputFromPrompt(projectNamePromptContent)

	projectTypePromptContent := utils.NewPromptContent(utils.PromptContentInput{
		Label: "⏳ Please select a project type.",
		ErrorMsg: "❌ You must select a project type",
	})

    items := []string{"NestJS", "NextJS", "React.js", "VueJs"}
	projectType := utils.GetSelectInputFromPrompt(projectTypePromptContent, items)
	fmt.Println(projectName)
	fmt.Println(projectType)
}



func init() {
	rootCmd.AddCommand(initCmd)
}
