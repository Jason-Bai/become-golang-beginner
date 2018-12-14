package main

import (
	"fmt"
	"unsafe"
)

func boolType() {
	a := true
	b := false
	fmt.Println("a: ", a, " b: ", b)
	c := a && b
	fmt.Println("a: ", a, " b: ", b, " c: ", c)
	d := a || b
	fmt.Println("a: ", a, " b: ", b, " c: ", c, " d: ", d)
}

func integerType() {
	var a int = 89
	b := 95
	fmt.Println("value of a is ", a, "and b is ", b)
}

func unsafeSizeOf() {
	var a int = 89
	b := 95
	fmt.Println("value of a is ", a, "and b is ", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is %T, size of b is %d\n", b, unsafe.Sizeof(b))
}

func floatingPointType() {
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum:", sum, "diff:", diff)
	no1, no2 := 56, 89
	fmt.Println("sum:", no1+no2, "diff:", no1-no2)
}

func complexType() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}

func stringType() {
	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("My name is ", name)
}

func typeConversion() {
	i := 55
	j := 67.8
	sum := i + int(j)
	fmt.Println("sum:", sum)
	k := float64(i)
	fmt.Println("k: ", k)
}

func main() {
	fmt.Println("1. Bool Type")
	boolType()
	fmt.Println("2. Integer Type")
	integerType()
	fmt.Println("3. unsafe.Sizeof")
	unsafeSizeOf()
	fmt.Println("4. Floating point types")
	floatingPointType()
	fmt.Println("4. Complex Type")
	complexType()
	fmt.Println("4. String Type")
	stringType()
	fmt.Println("5. Type Conversion")
	typeConversion()
}
