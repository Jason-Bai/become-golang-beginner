package main

import "fmt"

func simpleFor() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func breakFor() {
	for i := 1; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
}

func continueFor() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func onlyCondition() {
	i := 10
	for i < 10 {
		fmt.Printf("%d", i)
		i += 2
	}
}

func infiniteLoop() {
	for {
		fmt.Println("Hello world!")
	}
}

func main() {
	fmt.Println("1. simple for example")
	simpleFor()
	fmt.Println("2. break for")
	breakFor()
	fmt.Println("3. continue for")
	continueFor()
	fmt.Println("4. only condtion")
	onlyCondition()
	fmt.Println("5. infinite loop ")
	infiniteLoop()
}
