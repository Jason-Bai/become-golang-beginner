package main

import (
	"fmt"
	"unicode/utf8"
)

func whatIsAString() {
	name := "Hello World"
	fmt.Println(name)
}

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Printf("\n")
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Printf("\n")
}

func printChars2(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Printf("\n")
}

func accessIndividualBytesOfAString() {
	name := "Hello world"
	printBytes(name)
}

func accessIndividualCharsOfAString() {
	name := "Hello world"
	printChars(name)
}

func printUTF8Chars() {
	name := "Señor"
	printChars(name)
}

func printUTF8Chars2() {
	name := "Señor"
	printChars2(name)
}

func rangeOfString() {
	name := "Señor"
	for index, rune := range name {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func stringFromSlice() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
}

func stringFromSlice2() {
	byteSlice := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	fmt.Println(str)
}

func stringFromSlice3() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)
}

func lengthOfString() {
	word1 := "Señor"
	len1 := utf8.RuneCountInString(word1)
	fmt.Println("len1: ", len1)
	word2 := "Pets"
	len2 := utf8.RuneCountInString(word2)
	fmt.Println("len2: ", len2)
}

func mutateString() {
	h := "hello"
	fmt.Println("h: ", h)
	runes := []rune(h)
	runes[0] = 'H'
	fmt.Println("h: ", string(runes))
}

func main() {
	fmt.Println("1. what is a string")
	whatIsAString()
	fmt.Println("2. accessing individual bytes of a string")
	accessIndividualBytesOfAString()
	fmt.Println("2. accessing individual chars of a string")
	accessIndividualCharsOfAString()
	fmt.Println("3. print utf8 chars of string, but result is wrong")
	printUTF8Chars()
	fmt.Println("4. print utf8 chars of string")
	printUTF8Chars2()
	fmt.Println("5. range of a string")
	rangeOfString()
	fmt.Println("6. constructing string from slice of bytes")
	stringFromSlice()
	stringFromSlice2()
	fmt.Println("7. constructing a string from slice of runes")
	stringFromSlice3()
	fmt.Println("8. length of the string")
	lengthOfString()
	fmt.Println("9. mutate string")
	mutateString()
}
