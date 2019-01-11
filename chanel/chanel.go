package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan int{
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c\n", id, <-c)
		}
	}()
	return c
}
func chanDemo() {
	var channels [10] chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)

	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a'+i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A'+i
	}
	time.Sleep(time.Minute)
}

func bufferedChannel()  {
	c := make(chan int,3)
	c <-1
	c<-2
	c<-3
	c<-4

}

func channelClose()  {
	
}

func main() {
	//chanDemo()
	bufferedChannel()
}
