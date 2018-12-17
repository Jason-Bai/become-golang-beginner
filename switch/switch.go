package main

import "fmt"

func simpleSwitch() {
	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}
}

func defaultCase() {
	switch finger := 8; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("incorrect finger number")
	}
}

func multipleExpressions() {
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}

func expressionlessSwitch() {
	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	default:
		fmt.Println("num is greater than 100")
	}
}

func number() int {
	num := 15 * 5
	return num
}

func fallthroughSwitch() {
	switch num := number(); {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}

func main() {
	fmt.Println("1. simple switch")
	simpleSwitch()
	fmt.Println("2. default case")
	defaultCase()
	fmt.Println("3. multiple expressions")
	multipleExpressions()
	fmt.Println("4. expressionless switch")
	expressionlessSwitch()
	fmt.Println("4. fallthrough switch")
	fallthroughSwitch()
}
