package cmd

import (
	"fmt"

	"github.com/Onboardbase/obbkitv2/kit"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		gitUrl, err := cmd.Flags().GetString("from-git")

		if gitUrl == "" || err != nil {
			kit.Init()
			return nil
		}

		fmt.Println("Initializing from git repo: ", gitUrl)
		kit.InitFromGit(kit.InitFromGitInput{
			GitRepoURL: gitUrl,
		})
		return nil
	},
	Aliases: []string{"i"},
}

func init() {
	rootCmd.PersistentFlags().StringP("from-git", "g", "", "Initialize a project from a git repository")
	rootCmd.AddCommand(initCmd)
}
