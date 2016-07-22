package learn_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLearn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Learn Suite")
}
