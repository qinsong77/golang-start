package main

import (
	"fmt"
	"io"
	"strings"
)

func fibonacci() func() int  {
	a, b := 0, 1
	fmt.Printf("a = %d,b = %d\n",a,b)
	return func() int {
		a,b = b, a+b
		return  a
	}
}

type intGen func() int

func (g intGen) Read(p []byte)(n int, err error)  {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)
}
func main() {
	f := fibonacci()
	for i:= 0;i<10;i++ {
		fmt.Println(f())
	}
}
