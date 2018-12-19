package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func panicExample() {
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

func fullName2(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func deferWhilePanicking() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName2(&firstName, nil)
	fmt.Println("returned normally from main")
}

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName3(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recoverExample() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName3(&firstName, nil)
	fmt.Println("returned normally from main")
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	// go b()
	b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func recoverInSameGoroutine() {
	a()
	fmt.Println("normally returned from main")
}

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func c() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func runtimePanic() {
	c()
	fmt.Println("normally returned from main")
}

func r1() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}

func c1() {
	defer r1()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func stackTrack() {
	c1()
	fmt.Println("normally returned from main")
}

func main() {
	fmt.Println("1. panic example")
	// panicExample()
	fmt.Println("2. defer while panicking")
	// deferWhilePanicking()
	fmt.Println("3. recover example")
	// recoverExample()
	fmt.Println("4. recover in same goroutine")
	// recoverInSameGoroutine()
	fmt.Println("5. runtime panic")
	// runtimePanic()
	fmt.Println("6. stack trace in recover function")
	stackTrack()
}
