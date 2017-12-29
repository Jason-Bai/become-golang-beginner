package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://6brand.com")
	if err != nil {
		// handle error
		fmt.Println("retrieve error : ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read all error : ", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
