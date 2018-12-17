package main

import "fmt"

func declaration() {
	var a [3]int
	fmt.Println(a)
}

func assignSomeValue() {
	var a [3]int
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println(a)
}

func shortHandAArray() {
	a := [3]int{12, 78, 50}
	fmt.Println(a)
}

func noSpecialLength() {
	a := [...]int{12, 78, 50}
	fmt.Println(a)
}

func sizeIsPartOfType() {
	a := [3]int{5, 78, 8}
	var b [5]int
	// b = a //not possible since [3]int and [5]int are distinct types
	fmt.Println(a, b)
}

func arrayAreValueTypes() {
	// a := []string{"USA", "China", "India", "Germany", "France"} // this line will output different results
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}

func lengthOfArray() {
	a := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is ", len(a))
}

func iteratingArraysUsingRange() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d the elment of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("sum of all elements of a", sum)
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Printf("\n")
	}
}

func multiDimensionalArrays() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	printarray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)

}

func createASlice() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

func createASliceAnotherWay() {
	c := []int{6, 7, 8}
	fmt.Println(c)
}

func modifyingASlice() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before: ", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after: ", darr)
}

func reflectedTheArray() {
	numa := [3]int{78, 79, 80}
	nums1 := numa[:]
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
}

func lengthAndCapacityOfASlice() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
}

func reSlicedUptoASlice() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                        //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
}

func creatingASliceUsingMake() {
	i := make([]int, 5, 5)
	fmt.Println(i)
}

func appendingToASlice() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
}

func zeroValueOfASlice() {
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}
}

func appendOneSliceToAnother() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
}

func subtractOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func pointerVariable() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtractOne(nos)
	fmt.Println("slice after function call", nos)
}

func multiDimensionalSlices() {
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Printf("\n")
	}
}

func copyASlice() {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	fmt.Println(countriesCpy)
}

func main() {
	fmt.Println("1. declaration")
	declaration()
	fmt.Println("2. assign some value")
	assignSomeValue()
	fmt.Println("3. short hand a array")
	shortHandAArray()
	fmt.Println("4. no special length for array")
	noSpecialLength()
	fmt.Println("5. The size of the array is a part of the type.")
	sizeIsPartOfType()
	fmt.Println("6. arrays are value types")
	arrayAreValueTypes()
	fmt.Println("7. length of an array")
	lengthOfArray()
	fmt.Println("8. iterating array using range")
	iteratingArraysUsingRange()
	fmt.Println("9. multidimensional arrays")
	multiDimensionalArrays()
	fmt.Println("10. creating a slice")
	createASlice()
	fmt.Println("11. creating a slice another way")
	createASliceAnotherWay()
	fmt.Println("12. modifying a slice")
	modifyingASlice()
	fmt.Println("14. reflect the array")
	reflectedTheArray()
	fmt.Println("15. length and capacity of a slice")
	lengthAndCapacityOfASlice()
	fmt.Println("16. re-slicing upto a slice")
	reSlicedUptoASlice()
	fmt.Println("17. creating a slice using make")
	creatingASliceUsingMake()
	fmt.Println("18. appending to a slice")
	appendingToASlice()
	fmt.Println("19. The zero value of a slice type is nil")
	zeroValueOfASlice()
	fmt.Println("20. append one slice to another")
	appendOneSliceToAnother()
	fmt.Println("21. the pointer variable will refer to the same underlying array")
	pointerVariable()
	fmt.Println("22. multidimensional slices")
	multiDimensionalSlices()
	fmt.Println("23. copy a slice")
	copyASlice()
}
