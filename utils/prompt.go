package utils

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/manifoldco/promptui"
)

type promptContent struct {
	label    string
	errorMsg string
}

type PromptContentInput struct {
	Label    string
	ErrorMsg string
}

func NewPromptContent(input PromptContentInput) promptContent {
	return promptContent{input.Label, input.ErrorMsg}
}

func GetInputFromPrompt(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		justChars := regexp.MustCompile(`^[A-Za-z0-9]*$`).MatchString(input)
		if !justChars {
			return errors.New("❌ You must provide a project name with only letters and numbers")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | white  }} ",
		Confirm: "{{ . | cyan }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func GetSelectInputFromPrompt(pc promptContent, items []string) string {
	index := -1
	var result string
	var err error
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F449 {{ . | cyan }}",
		Selected: "\U0001F44D Project Type: {{ . | green | bold }} ",
	}
	for index < 0 {
		prompt := promptui.Select{
			Label:     pc.label,
			Items:     items,
			Templates: templates,
		}

		index, result, err = prompt.Run()
	}
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
