package graphparser_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGraphparser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphparser Suite")
}
