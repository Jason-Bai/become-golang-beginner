package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {
	fmt.Printf("State %s Country %s\n", a.state, a.country)
}

func pointerReceiver() {
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}
	// d2 = a
	d2 = &a
	d2.Describe()
}

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeaveLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeaveLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func implementingMultipleInterfaces() {
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanthan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}

	var s SalaryCalculator = e
	s.DisplaySalary()
	var l LeaveCalculator = e
	fmt.Println("\nLeaves left = ", l.CalculateLeaveLeft())
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

func embeddingInterfaces() {
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 50,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left = ", empOp.CalculateLeaveLeft())
}

type Describer2 interface {
	Describe()
}

func zeroValueOfInterface() {
	var d1 Describer2
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}
}

func main() {
	fmt.Println("1. pointer receiver")
	pointerReceiver()
	fmt.Println("2. implementing multiple interfaces")
	implementingMultipleInterfaces()
	fmt.Println("3. embedding interfaces")
	embeddingInterfaces()
	fmt.Println("4. zero value of interface")
	zeroValueOfInterface()
}
