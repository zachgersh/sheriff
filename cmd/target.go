package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(targetCmd)
}

var targetCmd = &cobra.Command{
	Use:   "target",
	Short: "UAA environment to target",
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(1)
	},
}
