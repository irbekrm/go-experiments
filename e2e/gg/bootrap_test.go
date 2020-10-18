package gg

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestGg(t *testing.T) {
	// This line connects Ginkgo with Gomega
	gomega.RegisterFailHandler(ginkgo.Fail)
	// Start the test suite
	ginkgo.RunSpecs(t, "Test gg")

}
