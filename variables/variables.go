package main

import "fmt"

func singleVariable() {
	var age int
	fmt.Println("my age is ", age)
	age = 29
	fmt.Println("my age is ", age)
	age = 54
	fmt.Println("my new age is ", age)
}

func variableWithInitialValue() {
	var age int = 29
	fmt.Println("my age is ", age)
}

func typeInference() {
	var age = 29
	fmt.Println("my age is ", age)
}

func multiVariableDeclaration() {
	var width, height int = 100, 50
	fmt.Println("width is ", width, " height is ", height)
}

func multiVariableDiffTypesDeclaration() {
	var (
		name   = "naveen"
		age    = 29
		height int
	)
	fmt.Println("my name is ", name, " age is ", age, " height is ", height)
}

func shortHandDeclaration() {
	name, age := "naveen", 29
	fmt.Println("my name is ", name, " age is ", age)
}

func main() {
	fmt.Println("1. Declaring a single variable")
	singleVariable()
	fmt.Println("2. Declaring a variable with initial value")
	variableWithInitialValue()
	fmt.Println("3. Type inference")
	typeInference()
	fmt.Println("4. Multiple variable declaration")
	multiVariableDeclaration()
	fmt.Println("5. Multiple variable types declaration")
	multiVariableDiffTypesDeclaration()
	fmt.Println("5. Shord hand declaration")
	shortHandDeclaration()
}
