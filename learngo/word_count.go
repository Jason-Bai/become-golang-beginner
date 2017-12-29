package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "1你好12dfsdf"
	fmt.Printf("String %s Length: %d, Runes: %d\n", str, len(str), utf8.RuneCount([]byte(str)))
}
