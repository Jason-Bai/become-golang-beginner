package main

import "fmt"

// Employee a structure of employee
type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func exampleMethod() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}

	emp1.displaySalary()
}

// Employee2 a structure of employee2
type Employee2 struct {
	name string
	age  int
}

func (e Employee2) changeName(newName string) {
	e.name = newName
}

func (e *Employee2) changeAge(newAge int) {
	e.age = newAge
}

func receiversVs() {
	e := Employee2{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	//(&e).changeAge(51)
	e.changeAge(51)
	fmt.Printf("\nEmployee age after change: %d\n", e.age)
}

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

func methodsOfAnonymousFields() {
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "Califonia",
		},
	}

	p.fullAddress()
}

type rectangle struct {
	length, width int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func valueReceiverMethod() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()
	p := &r
	p.area()
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output: ", 2*(r.length+r.width))
}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output: ", 2*(r.length+r.width))
}

func pointerReceiverMethod() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	p := &r
	perimeter(p)
	p.perimeter()
	r.perimeter()
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func methodsOnNonStructTypes() {
	num1 := myInt(5)
	num2 := myInt(0)
	sum := num1.add(num2)
	fmt.Println("sum is ", sum)
}

func main() {
	fmt.Println("1. example method")
	exampleMethod()
	fmt.Println("2. pointer receivers vs value receivers")
	receiversVs()
	fmt.Println("3. methods of anonymous fields")
	methodsOfAnonymousFields()
	fmt.Println("4. value receivers in methods vs value arguments in functions ")
	valueReceiverMethod()
	fmt.Println("5. pointer receivers in methods vs pointer arguments in functions")
	pointerReceiverMethod()
	fmt.Println("6. Methods on non struct types")
	methodsOnNonStructTypes()
}
