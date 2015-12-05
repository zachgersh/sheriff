package mod

import "github.com/spf13/cobra"

type target struct {
}

type fs interface {
}

func NewTargetModule(fs fs, configDir string) func(*cobra.Command, []string) {
	return target{}.Save
}

func (target) Save(command *cobra.Command, args []string) {
}
