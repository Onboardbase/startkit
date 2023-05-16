package utils

import (
	"fmt"
	"os"
	"path"
)

func GetProjectPath(folderName string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	projectDir := path.Join(currentDir, folderName)
	return projectDir
}
