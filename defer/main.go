package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writeFile("fibonacci.txt")
}
func writeFile(filename string)  {
	file,err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer  writer.Flush()
	f := fibonacci()
	for i := 0;i <20; i++{
		fmt.Fprintln(writer,f())
	}

}
func fibonacci() func() int  {
	a, b := 0, 1
	return func() int {
		a,b = b, a+b
		return  a
	}
}