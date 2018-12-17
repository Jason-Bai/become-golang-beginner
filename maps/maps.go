package main

import "fmt"

func createAMap() {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil, Going to make one.")
		personSalary = make(map[string]int)
	}
	personSalary["jason"] = 100

	for i, v := range personSalary {
		fmt.Println("i : ", i, " v : ", v)
	}
}

func addItems2Map() {
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents: ", personSalary)
}

func initializeAMap() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents: ", personSalary)
}

func accessItemsOfAMap() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	employee := "jamie"
	fmt.Println("Salary of ", employee, " is ", personSalary[employee])
}

func notPresent() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])
	fmt.Println("Salary of joe is", personSalary["joe"])
}

func checkKeyInMap() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	newEmp := "joe"
	value, ok := personSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}
}

func rangeAMap() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("All items of a map")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
}

func deleteItems() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("map before deletion", personSalary)
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)
}

func lengthOfTheMap() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("length is", len(personSalary))
}

func mapIsReferenceType() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)
}

func main() {
	fmt.Println("1. create a map")
	createAMap()
	fmt.Println("2. add items to a map")
	addItems2Map()
	fmt.Println("3. initialize a map")
	initializeAMap()
	fmt.Println("4. access items of a map")
	accessItemsOfAMap()
	fmt.Println("5. access not present items of a map")
	notPresent()
	fmt.Println("6. find out whether a particular key is present in a map")
	checkKeyInMap()
	fmt.Println("7. range a map")
	rangeAMap()
	fmt.Println("8. delete items from a map")
	deleteItems()
	fmt.Println("9. length of the map")
	lengthOfTheMap()
	fmt.Println("10. maps are reference types")
	mapIsReferenceType()
}
