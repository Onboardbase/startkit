package utils

import (
	"os"
	"os/exec"
)

type customStdIn struct {
}
func (c customStdIn) Read(buf []byte) (n int, err error) {
	return os.Stdin.Read(buf)
}

type RunShellCommandInput struct {
	ShellToUse string
	Command string
	Directory string
}

func RunShellCommand(input RunShellCommandInput) (err error) {
	cmd := exec.Command(input.ShellToUse, "-c", input.Command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if input.Directory != "" {
		cmd.Dir = input.Directory
	}
	err = cmd.Run()
	return
}