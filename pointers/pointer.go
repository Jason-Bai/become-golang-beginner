package main

import "fmt"

func declaringAPointer() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is ", a)
}

func zeroValueOfAPointer() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is ", b)
		b = &a
		fmt.Println("b after initialization is ", b)
	}
}

func dereferencingAPointer() {
	b := 255
	a := &b
	fmt.Println("address of b is ", a)
	fmt.Println("value of b is ", *a)
	*a++
	fmt.Println("new value of b is", b)
}

func change(val *int) {
	*val = 55
}

func passingPointer2AFunction() {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is", a)
}

func modify(arr *[3]int) {
	(*arr)[0] = 90
}

func doNotPassAPonterToAnArrayAsAArgumentToAFunction() {
	a := [3]int{89, 90, 91}
	modify(&a)
	fmt.Println(a)
}

func modify2(sls []int) {
	sls[0] = 90
}

func passSliceAsAnArrayArgumentToAFunction() {
	a := [3]int{89, 90, 91}
	modify2(a[:])
	fmt.Println(a)
}

func main() {
	fmt.Println("1. Declaring pointers")
	declaringAPointer()
	fmt.Println("2. zero value of a pointers")
	zeroValueOfAPointer()
	fmt.Println("3. dereferencing a pointer")
	dereferencingAPointer()
	fmt.Println("4. passing pointer to a function")
	passingPointer2AFunction()
	fmt.Println("5. do not pass a pointer to an array as a argument to a function")
	doNotPassAPonterToAnArrayAsAArgumentToAFunction()
	fmt.Println("6. pass a pointer to a slice as an array argument to a function")
	passSliceAsAnArrayArgumentToAFunction()
}
