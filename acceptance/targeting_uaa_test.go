package acceptance_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/naoina/toml"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type config struct {
	Targets map[string]target
}

type target struct {
	Active bool
	URL    string
}

var _ = Describe("Targeting UAA", func() {
	var homeDir = os.Getenv("HOME")

	It("saves a target", func() {
		uaa_url := "uaa.run.pivotal.io"
		command := exec.Command(sheriffPath, "target", uaa_url)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ShouldNot(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))

		Eventually(session.Out).Should(gbytes.Say(fmt.Sprintf("Targeted: %s", uaa_url)))

		raw, err := ioutil.ReadFile(path.Join(homeDir, ".sheriff.toml"))
		Expect(err).ShouldNot(HaveOccurred())

		var conf config

		err = toml.Unmarshal(raw, &conf)
		Expect(err).ShouldNot(HaveOccurred())

		Expect(conf.Targets).To(HaveKeyWithValue(`targets."run.pivotal"`, target{Active: true, URL: "uaa.run.pivotal.io"}))
	})
})
