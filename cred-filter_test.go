package main_test

import (
	// . "github.com/geramirez/concourse-filter"

	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func runBinary(stdin string, env []string) (string, error) {
	cmd := exec.Cmd{
		Path:  "./cred-filter.exe",
		Stdin: strings.NewReader(stdin),
		Env:   env,
	}
	output, err := cmd.Output()
	return string(output), err
}

var _ = Describe("CredFilter", func() {

	Context("No sensitive credentials available", func() {
		It("outputs as is", func() {
			env := []string{}
			output, err := runBinary("boring text", env)
			Expect(err).To(BeNil())
			Expect(output).To(Equal("boring text\n"))
		})
	})
	Context("Sensitive credentials available", func() {
		It("filters out those credentials", func() {
			env := []string{"SECRET=secret"}
			output, err := runBinary("super secret info", env)
			Expect(err).To(BeNil())
			Expect(output).To(Equal("super [redacted] info\n"))
		})
	})
})