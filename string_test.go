package learn_test

import (
	"learn"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String", func() {
	It("convert to []byte", func() {
		bs, _ := learn.StringToByteSlice("hello")
		Expect(bs).Should(Equal([]byte{'h', 'e', 'l', 'l', 'o'}))

		Expect(learn.StringToByteSlice2("hello")).Should(Equal([]byte{'h', 'e', 'l', 'l', 'o'}))
	})

	It("convert any object to string", func() {
		Expect(learn.String(10)).Should(Equal("10"))
		Expect(learn.String(10.10)).Should(Equal("10.1"))
		Expect(learn.String(true)).Should(Equal("true"))
	})

})

var _ = Describe("ByteSlice", func() {
	It("convert to string", func() {

		Expect(learn.ByteSlice2String([]byte{'h', 'e', 'l', 'l', 'o'})).Should(Equal("hello"))
	})

})
