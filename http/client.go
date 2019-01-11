package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, error := http.Get("http://www.imooc.com")
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	response,error := httputil.DumpResponse(resp,true)
	if error != nil {
		panic(error)
	}
	fmt.Printf("%s",response)
}
