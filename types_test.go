package learn_test

import (
	"fmt"
	. "learn"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Types", func() {

	It("Str2Int64", func() {
		var r int64 = 0
		a, err := Str2Int64("")
		fmt.Println(err)
		Expect(a).Should(Equal(r))
	})

})
