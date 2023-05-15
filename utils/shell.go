package utils

import (
	"os"
	"os/exec"
)

type RunShellCommandInput struct {
	ShellToUse string
	Command string
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
	err = cmd.Run()
	return
}