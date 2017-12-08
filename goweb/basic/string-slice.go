package main

import "fmt"

func main() {
	s := "hello"
	s2 := "c" + s[1:]
	fmt.Printf("%s\n", s2)
}
