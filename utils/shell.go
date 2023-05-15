package utils

import (
	"fmt"
	"os"
	"os/exec"
)

type RunShellCommandInput struct {
	ShellToUse       string
	Command          string
	DirectoryToRunIn string
}

func RunShellCommand(input RunShellCommandInput) (err error) {
	cmd := exec.Command(input.ShellToUse, "-c", input.Command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if input.DirectoryToRunIn != "" {
		cmd.Dir = input.DirectoryToRunIn
	}
	if err = cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
		os.Exit(1)
	}
	return
}
