package main

import "fmt"

func main() {
	arr1 := [2] string { "blue", "red"}
	arr2 := [2] string {}
	arr2 = arr1
	arr1[1] = "yellow"
	fmt.Println(&arr1[0])
	fmt.Println(&arr1[1])
	fmt.Println(arr1)
	fmt.Println(arr2)

	arr3 := [2] *string{new(string), new(string)}
	*arr3[0] = "blue"
	*arr3[1] = "red"
	arr4 := arr3
	fmt.Println(arr3)
	*arr3[1] = "yellow"
	for _,v := range arr3 {
		fmt.Println(v)
	}
	for v := range arr4 {
		fmt.Println(v)
	}
	fmt.Println(arr3)
	fmt.Println(arr4)

	slice := [] int {1,2,3,4,5}
	fmt.Println(slice)
	newSlice := slice[1:3]
	newSlice[1] = 100
	fmt.Println(newSlice)
	fmt.Println(slice)
	test(slice)
	fmt.Println(slice)
}

func test(sli []int)   {
	sli[0] = 1212
}