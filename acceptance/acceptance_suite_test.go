package acceptance_test

import (
	"io/ioutil"
	"os"

	"testing"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}

var (
	sheriffPath, tempPath string
)

var _ = BeforeSuite(func() {
	var err error
	sheriffPath, err = gexec.Build("github.com/zachgersh/sheriff")
	Expect(err).ShouldNot(HaveOccurred())

	tempPath, err = ioutil.TempDir("", "sheriff")
	Expect(err).ShouldNot(HaveOccurred())

	os.Setenv("HOME", tempPath)
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()

	err := os.Unsetenv("HOME")
	Expect(err).ShouldNot(HaveOccurred())

	err = os.RemoveAll(tempPath)
	Expect(err).ShouldNot(HaveOccurred())
})
