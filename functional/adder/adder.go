package main

import "fmt"

func adder() func(int) int  {
	sum := 0
	return func(v int) int {
		sum += v
		fmt.Printf("sum is %d\n", sum)
		return sum
	}
}
func main() {
	a := adder()
	for i := 0;i<10;i++ {
		fmt.Printf("0 + 1 + ... %d = %d\n",i,a(i))
	}
}
