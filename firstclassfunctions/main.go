package main

import "fmt"

func anonymousFunctions() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T\n", a)
}

func anonymousFunctions2() {
	func() {
		fmt.Println("hello world first class function")
	}()
}

func anonymousFunctions3() {
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
}

type add func(a, b int) int

func userDefinedFunctionTypes() {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum: ", s)
}

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func highOrderFunctions() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}

func hof() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func highOrderFunctions2() {
	s := hof()
	fmt.Println(s(60, 7))
}

func closuresExample() {
	a := 5
	func() {
		fmt.Println("a = ", a)
	}()
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t += " " + b
		return t
	}
	return c
}

func closuresExample2() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

type student struct {
	firstName, lastName, grade, country string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func filterExample() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func iMapExample() {
	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}

func main() {
	fmt.Println("1. Anonymous functions")
	anonymousFunctions()
	fmt.Println("2. Anonymous functions 2")
	anonymousFunctions2()
	fmt.Println("3. Anonymous functions 3")
	anonymousFunctions3()
	fmt.Println("4. User defined function types")
	userDefinedFunctionTypes()
	fmt.Println("5. high order functions")
	highOrderFunctions()
	fmt.Println("6. high order functions2")
	highOrderFunctions2()
	fmt.Println("7. closures")
	closuresExample()
	fmt.Println("8. closures 2")
	closuresExample2()
	fmt.Println("9. filter example")
	filterExample()
	fmt.Println("10. iMap example")
	iMapExample()
}
