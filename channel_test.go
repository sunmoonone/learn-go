package learn_test

import (
	"fmt"
	. "learn"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func sleep(mili int64) {
	time.Sleep(time.Duration(mili) * time.Millisecond)
}

var _ = Describe("Goroutine", func() {

	PContext("share variable between routines", func() {

		Specify("Two routines can access to the same variable", func() {

			var i int = 0
			sig := &i
			go func(signal *int) {

				for {

					time.Sleep(10 * time.Millisecond)
					if *signal == -1 {
						fmt.Println("exit routine")
						*signal = 10
						break
					}
				}

			}(sig)

			go func() {
				time.Sleep(50 * time.Millisecond)
				*sig = -1
			}()

			time.Sleep(100 * time.Millisecond)
			Expect(*sig).Should(Equal(10))
		})

	})

})

var _ = Describe("Channel", func() {

	/*Context("Serve", func() {

		queue := make(chan *Request)
		result := make([]string, 5)

		XSpecify("run forever", func(done Done) {
			put 5 requests
			go func() {
				for i := 0; i < 5; i++ {
					queue <- &Request{string(i)}
					time.Sleep(100 * time.Millisecond)
				}

			}()

			Serve(queue, result)
			Expect(len(result)).Should(Equal(5))

			close(done)
		}, 5)

	})
	*/

	XContext("Serve2", func() {

		queue := make(chan *Request)
		resultV := make([]string, 0)
		result := &resultV

		/*
			* the following two lines may cause a panic of memory space error
			* because the pointer points to nil address
			var quit *int
			quit = 0
		*/
		var quitV int = 0
		quit := &quitV

		Specify("quit on variable value", func(done Done) {

			go func() {
				for i := 0; i < 5; i++ {
					queue <- &Request{fmt.Sprintf("%d", i)}
					sleep(10)
				}

			}()

			go Serve2(5, queue, result, quit)

			sleep(2000)
			*quit = 1
			for _, r := range *result {
				fmt.Printf("\n result %s", r)
			}

			Expect(len(*result)).Should(Equal(5))

			close(done)
		}, 5)

	})

	XContext("Serve3", func() {

		queue := make(chan *Request)
		quit := make(chan bool)
		result := make(chan string, 5)

		Specify("quit on variable value", func(done Done) {

			go func() {
				for i := 0; i < 5; i++ {
					queue <- &Request{fmt.Sprintf("%d", i)}
					sleep(10)
				}

			}()

			go Serve3(5, queue, result, quit)
			quit <- true

			sleep(2000)
			Expect(len(result)).Should(Equal(5))
			i := len(result)
			for r := range result {
				fmt.Printf("\n result %s", r)
				i--
				if i == 1 {
					break
				}
			}

			close(done)
		}, 5)

	})
})
