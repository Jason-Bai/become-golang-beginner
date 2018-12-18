package main

import (
	"fmt"

	"github.com/Jason-Bai/golang_codes/structsInsteadOfClasses/oop/employee"
)

func structsInsteadOfClasses() {
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}

func main() {
	fmt.Println("1. Structs Instead of Classes")
	structsInsteadOfClasses()
}
