package gg

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Basic syntax-", func() {
	ginkgo.Context("gomega matchers-", func() {
		ginkgo.It("ExpectWithOffset can be used in helper functions to 'step up' call stack", func() {
			mustBeBar("barr")
		})
	})
})

func mustBeBar(s string) {
	// If this line errors, print the line number one level up - when this function was called
	gomega.ExpectWithOffset(1, s).To(gomega.Equal("bar"))
}
