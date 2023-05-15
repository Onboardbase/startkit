package cmd

import (
	"github.com/Onboardbase/obbkitv2/kit"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "I",
	Run: func(cmd *cobra.Command, args []string) {
		kit.CreateProject()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
