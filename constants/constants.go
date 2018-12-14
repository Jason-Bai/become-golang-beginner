package main

import (
	"fmt"
	"math"
)

func cannotAssignAgain() {
	const a = 55
	// a = 89 // throw a error
	fmt.Println("a: ", a)
}

func knownAtCompileTime() {
	var a = math.Sqrt(4)
	// const b = math.Sqrt(4) // throw a error
	fmt.Println("a: ", a)
}

const hello = "HelloWorld"

func unTypedString() {
	var name = "Sam"
	fmt.Printf("type of name is %T and name value is %v\n", name, name)
	fmt.Printf("type of hello is %T and hello value is %v\n", hello, hello)
}

func stringMixingType() {
	var defaultName = "Sam"
	type myString string
	var customName myString
	// customName = defaultName // can't convert
	fmt.Println("defaultName: ", defaultName, " customName: ", customName)
}

func boolMixingType() {
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       // allowed
	var customBool myBool = trueConst // allowed
	// defaultBool = customBool // not allowed
	fmt.Println("defaultBool: ", defaultBool, " customBool: ", customBool)
}

func numberMixingType() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}

func numericExpressions() {
	var a = 5.9 / 8
	fmt.Printf("a's type %T value %v\n", a, a)
}

func main() {
	fmt.Println("1. Can't Assign")
	cannotAssignAgain()
	fmt.Println("2. Constant should be known at compile time")
	knownAtCompileTime()
	fmt.Println("3. Untyped String")
	unTypedString()
	fmt.Println("4. String Mixing Types")
	stringMixingType()
	fmt.Println("5. Bool Mixing Types")
	boolMixingType()
	fmt.Println("6. Number Mixing Types")
	numberMixingType()
	fmt.Println("7. Numeric Expressions")
	numericExpressions()
}
