package kit

import (
	"fmt"
	"os"
	"path"

	"github.com/Onboardbase/obbkitv2/utils"
)

func initNestJsProject(projectName string) {
	fmt.Println("\t \U000023F3 Creating NestJS project...")
	cmd := fmt.Sprintf("npx @nestjs/cli new %s", projectName)
	 err := utils.RunShellCommand(utils.RunShellCommandInput{
		ShellToUse: "bash",
		Command:    cmd,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd = "onboardbase auth -c \"yarn start\""
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	projectName = utils.NormalizeToKebabOrSnakeCase(projectName)
	fmt.Println("projectName", projectName)
	projectDir := path.Join(currentDir, projectName)
	err = utils.RunShellCommand(utils.RunShellCommandInput{
		ShellToUse: "bash",
		Command:    cmd,
		Directory:  projectDir,
	})
	if err != nil {
		fmt.Println(err)
	}
}