package learn_test

import (
	"learn"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func InstanceOfDirection(v interface{}) bool {
	_, ok := v.(learn.Direction)
	return ok
}

var _ = Describe("Const", func() {
	It("can use a custom type", func() {
		Expect(InstanceOfDirection(learn.UP)).Should(BeTrue())
		Expect(InstanceOfDirection(learn.DOWN)).Should(BeTrue())
	})

	It("can use iota with custom type", func() {
		Expect(int(learn.UP)).Should(Equal(0))
		Expect(int(learn.DOWN)).Should(Equal(1))
		Expect(int(learn.RIGHT)).Should(Equal(3))
	})
})

var _ = Describe("Point", func() {
	p := learn.Point{0, 0}

	It("move up", func() {
		var _ = p.Move(learn.UP, 1)
		Expect(p.Y).Should(Equal(-1))
	})

	It("move down", func() {
		var _ = p.Move(learn.DOWN, 1)
		Expect(p.Y).Should(Equal(0))
	})

	It("move left", func() {
		var _ = p.Move(learn.LEFT, 1)
		Expect(p.X).Should(Equal(-1))
	})

	It("move right", func() {
		var _ = p.Move(learn.RIGHT, 2)
		Expect(p.X).Should(Equal(1))
	})
})

var _ = Describe("Int Pointer", func() {

	It("usage", func() {
		var i int = 0
		var pi *int = &i
		*pi++
		Expect(i).Should(Equal(1))
		Expect(pi).ShouldNot(Equal(i))

		pi2 := &i
		*pi2 = 2
		Expect(i).Should(Equal(2))

	})

	/*It("assign int to *int should fail", func() {
		i := 0
		pi := &i
		var f float32
		pi = f

		err := errors.New("")
		var invalidAssignment = func() {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.ErrorF("%v", r)
				}
			}()
			pi = 3
		}
		invalidAssignment()
		Expect(err.Error()).ShouldNot(Equal(""))

	})*/

})
