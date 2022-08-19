package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestParseArguments(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ParseArguments Test Suite")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var _ = Describe("argument parser", func() {
	It("should return a simple Argument", func() {
		Expect(ParseArguments("--help")).To(Equal(2))
	})
})
