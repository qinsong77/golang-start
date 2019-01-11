package main

import (
	"fmt"
	"time"
	"./src"
)

type myTime struct {
	time.Time
}

func (t myTime)first3Chars()string  {
	return  t.Time.String()[0:3]
}

type Retriever interface {
	Get(url string) string
}

func download(r Retriever)  string {
	return r.Get("http://www.baidu.com")
}
func main()  {
	//fmt.Println("main is running")
	//src.Test("2332")
	////src.Testinterface()
	//m := myTime{time.Now()}
	//fmt.Println("Full time now :",m.String())
	//fmt.Println("First 3 chars:",m.first3Chars())
	var r Retriever = src.Retriever{}
	fmt.Println(download(r))

}
