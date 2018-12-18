package main

import "fmt"

// VowelsFinder a interface of VowelsFinder
type VowelsFinder interface {
	FindVowels() []rune
}

// MyString new string
type MyString string

// FindVowels implemented FindVowels method
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func declaringAndImplementingAnInterface() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c\n", v.FindVowels())
}

// SalaryCalculator the interface of SalaryCalculator
type SalaryCalculator interface {
	CalculateSalary() int
}

// Permanent the struct of Permanent
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

// Contract the struct of Contract
type Contract struct {
	empId    int
	basicpay int
}

// CalculateSalary the method of Permanent
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// CalculateSalary the method of Contract
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d\n", expense)
}

func practicalUseOfInterface() {
	pemp1 := Permanent{1, 5000, 2}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}

type Tester interface {
	Test()
}

type MyFloat float64

func (m MyFloat) Test() {
	fmt.Println(m)
}

func describe(t Tester) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func interfaceInternalRepresentation() {
	var t Tester
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Test()
}

func describe2(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func emptyInterface() {
	s := "Hello world!"
	describe2(s)
	i := 55
	describe2(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe2(strt)
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func typeAssertion() {
	var s interface{} = 56
	assert(s)
	var a interface{} = "Steven Paul"
	assert(a)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am a int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func typeSwitch() {
	findType("Naveen")
	findType(77)
	findType(89.89)
}

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

func findType2(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func typeSwitch2() {
	findType2("Naveen")
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType2(p)
}

func main() {
	fmt.Println("1. Declaring and implementing an interface")
	declaringAndImplementingAnInterface()
	fmt.Println("2. practical use of interface")
	practicalUseOfInterface()
	fmt.Println("3. interface internal representation")
	interfaceInternalRepresentation()
	fmt.Println("4. empty interface")
	emptyInterface()
	fmt.Println("5. Type Assertion")
	typeAssertion()
	fmt.Println("6. Type Switch")
	typeSwitch()
	typeSwitch2()
}
