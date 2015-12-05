package mod_test

import (
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/zachgersh/sheriff/mod"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Target", func() {
	Context("when the target can be reached", func() {
		Context("when the file does not exist", func() {
			It("creates the .sheriff.toml with the target info", func() {
				fs := &afero.MemMapFs{}
				fs.Mkdir("/tmp", 0755)

				command := &cobra.Command{}
				args := []string{"some-uaa-url"}
				mod := mod.NewTargetModule(fs, "/tmp")

				mod(command, args)

				_, err := os.Stat("/tmp/.sheriff.toml")
				Expect(os.IsNotExist(err)).To(BeFalse())
			})
		})
	})
})
