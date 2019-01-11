package main

import (
	"fmt"
)

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		if id == 7 {
			done <- false
		}else {
			done <- true
		}
		c <- id
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}
func chanDemo() {
	var workers [10]  worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for _, worker := range workers {
		if <-worker.done {
			fmt.Printf("worker %d done\n", <-worker.in)
		}else {
			fmt.Printf("worker %d not done\n", <-worker.in)
		}
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	for _, worker := range workers {
		if <-worker.done {
			fmt.Printf("worker %d done\n", <-worker.in)
		}else {
			fmt.Printf("worker %d not done\n", <-worker.in)
		}
	}
}

func main() {
	chanDemo()
}
