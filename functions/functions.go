package main

import "fmt"

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func declarationAFunction() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is ", totalPrice)
}

func singleReturnValue(a, b int) int {
	var c = a + b
	return c
}

func multipleReturnValues(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func namedReturnValues(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func blankIdentifier() {
	area, _ := namedReturnValues(10.8, 5.6)
	fmt.Println("area: ", area)
}

func main() {
	fmt.Println("1. Declaration a function")
	declarationAFunction()
	fmt.Println("2. Single return  value")
	add := singleReturnValue(1, 2)
	fmt.Println("add: ", add)
	fmt.Println("3. Mutiple return values")
	area, perimeter := multipleReturnValues(10.8, 5.6)
	fmt.Println("area: ", area, " perimeter: ", perimeter)
	fmt.Println("4. Named return values")
	area, perimeter = namedReturnValues(10.8, 5.6)
	fmt.Println("area: ", area, " perimeter: ", perimeter)
	fmt.Println("5. Blank Identifier")
	blankIdentifier()
}
