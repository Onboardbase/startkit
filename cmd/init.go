package cmd

import (
	"github.com/Onboardbase/obbkit/kit"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize a new project",
	Aliases: []string{"i"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return handleInit(cmd)
	},
}

func init() {
	rootCmd.PersistentFlags().StringP(FROM_GIT_FLAG, "g", "", "Initialize a project from a git repository")
	rootCmd.AddCommand(initCmd)
}

func handleInit(cmd *cobra.Command) error {
	gitUrl, err := cmd.Flags().GetString(FROM_GIT_FLAG)
	if err == nil && gitUrl != "" {
		kit.InitFromGit(kit.InitFromGitInput{
			GitRepoURL: gitUrl,
		})
		return nil
	}
	kit.Init()
	return nil
}
