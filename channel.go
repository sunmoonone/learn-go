package learn

import (
	"fmt"
)

type Request struct {
	Cmd string
}

func process(r *Request, result *[]string) {
	*result = append(*result, r.Cmd)
	fmt.Printf("\nprocessing request:%s", r.Cmd)
}

func Serve(concurrencyCount int, queue chan *Request, result *[]string) {
	var sem = make(chan int, concurrencyCount)
	for req := range queue {
		sem <- 1
		go func(req *Request) {
			process(req, result)
			<-sem
		}(req)
	}
}

//======================

func Serve2(concurrencyCount int, queue chan *Request, result *[]string, quit *int) {
	var sem = make(chan int, concurrencyCount)
	for req := range queue {

		sem <- 1
		go func(req *Request) {
			process(req, result)
			<-sem
		}(req)

		if *quit == 1 {
			break
		}
	}
}

//======================
func process3(req *Request, result chan string) {
	result <- req.Cmd
	fmt.Printf("\nprocessing request:%s", req.Cmd)
}

func handle(queue chan *Request, result chan string) {
	for r := range queue {
		process3(r, result)
	}
}

func Serve3(concurrencyCount int, clientRequests chan *Request, result chan string, quit chan bool) {
	// Start handlers
	for i := 0; i < concurrencyCount; i++ {
		go handle(clientRequests, result)
	}
	<-quit // Wait to be told to exit.
}
