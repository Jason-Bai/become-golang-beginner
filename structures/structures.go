package main

import (
	"fmt"

	"github.com/Jason-Bai/golang_codes/structures/computer"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func creatingNamedStructures() {
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	emp2 := Employee{
		"Thomas",
		"Paul",
		29,
		800,
	}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
}

func creattingAnonymousStructures() {
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       29,
		salary:    5000,
	}
	fmt.Println("Employee 3", emp3)
}

func zeroValueOfAStructure() {
	var emp4 Employee
	fmt.Println("Employee 4", emp4)
}

func ignoreFields() {
	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("Employee 5", emp5)
}

func accessInDividualFieldsOfAStructure() {
	emp6 := Employee{
		"Sam",
		"Anderson",
		55,
		6000,
	}

	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age: ", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)
}

func assignValueOfAStructure() {
	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7: ", emp7)
}

func pointersToAStruct() {
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	// fmt.Println("First Name: ", (*emp8).firstName)
	fmt.Println("First Name: ", emp8.firstName)
	// fmt.Println("Age: ", (*emp8).age)
	fmt.Println("Age: ", emp8.age)
}

// Person a structure of person
type Person struct {
	string
	int
}

func anonymousFields() {
	p := Person{"Naveen", 50}
	fmt.Println(p)
}

func anonymousFields2() {
	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)
}

// Address a structure of address
type Address struct {
	city, state string
}

// Person2 a structure of person2 with address
type Person2 struct {
	name    string
	age     int
	address Address
}

func nestedStructs() {
	var p Person2
	p.name = "Naveen"
	p.age = 50
	p.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("City: ", p.address.city)
	fmt.Println("State: ", p.address.state)
}

func promotedFields() {
	type Person struct {
		name string
		age  int
		Address
	}

	var p Person

	p.name = "Naveen"
	p.age = 50
	p.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("City: ", p.city)
	fmt.Println("State: ", p.state)
}

func exportedStructsFields() {
	var spec computer.Spec
	spec.Marker = "apple"
	spec.Price = 50000
	fmt.Println("Spec: ", spec)
}

type name struct {
	firstName, lastName string
}

func structsEquality() {
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and nam4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}

type image struct {
	data map[int]int
}

func structsEquality2() {
	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}
	/*
		if image1 == image2 {
			fmt.Println("image1 and image2 are equal")
		}
	*/
	fmt.Println("image1: ", image1)
	fmt.Println("image2: ", image2)
}

func main() {
	fmt.Println("1. creating named structures ")
	creatingNamedStructures()
	fmt.Println("2. creating anonymous structures ")
	creattingAnonymousStructures()
	fmt.Println("3. zero value of structure")
	zeroValueOfAStructure()
	fmt.Println("4. specify values for some fields and ignore the rest")
	ignoreFields()
	fmt.Println("5. accessing individual fields of a struct")
	accessInDividualFieldsOfAStructure()
	fmt.Println("6. create a zero struct and then assign values")
	assignValueOfAStructure()
	fmt.Println("7. pointers to a structure")
	pointersToAStruct()
	fmt.Println("8. anonymous fields")
	anonymousFields()
	anonymousFields2()
	fmt.Println("9. nested strcuture")
	nestedStructs()
	fmt.Println("10. exported structs and fields")
	exportedStructsFields()
	fmt.Println("11. structs equality")
	structsEquality()
	fmt.Println("12. structs equality2")
	structsEquality2()
}
