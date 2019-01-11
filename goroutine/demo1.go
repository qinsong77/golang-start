package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutine")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char ++ {
				time.Sleep(1009)
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char ++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
