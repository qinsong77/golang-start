package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/file/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL.Path)
		path := request.URL.Path[len("/file/"):]
		file,err := os.Open(path)
		if err != nil {
			http.Error(
				writer,
				err.Error(),
				http.StatusNotFound)
			return
		}
		defer file.Close()

		all,err :=ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		writer.Write(all)
	})
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic(err)
	}
}
