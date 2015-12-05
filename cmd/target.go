package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zachgersh/sheriff/mod"
)

func init() {
	RootCmd.AddCommand(targetCmd)
}

var targetCmd = &cobra.Command{
	Use:   "target",
	Short: "UAA environment to target",
	Run:   mod.NewTargetModule(nil, nil),
}
