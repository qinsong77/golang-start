package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Microsecond)
			out <- i
			i ++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	worker := createWorker(0)
	n := 0
	hasValue := false
	for {
		var activeWorker chan<- int
		if hasValue {
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
			//fmt.Println("Received from c1:", n)
		case n = <-c2:
			hasValue = true
		//fmt.Println("Received from c2:", n)
		case activeWorker <- n:
			hasValue = false
		}
	}
}
