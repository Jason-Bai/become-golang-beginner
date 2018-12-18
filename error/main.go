package main

import (
	"fmt"
	"os"
)

func errorExample() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name())
}

func errorExample2() {
	f, err := os.Open("/test.txt")

	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

func main() {
	fmt.Println("1. error example")
	errorExample()
	fmt.Println("2. error example2")
	errorExample2()
}
