package main

import "fmt"

func simpleIfElse() {
	num := 10
	if num%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}

func statementCondition() {
	if num := 10; num%2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, " is odd")
	}
}

func elseIf() {
	num := 99
	if num <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}
}

func main() {
	fmt.Println("1. simple if-else")
	simpleIfElse()
	fmt.Println("2. statement condition")
	statementCondition()
	fmt.Println("3. else if")
	elseIf()
}
