package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func startAGoroutine() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("startAGoroutine")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Microsecond)
		fmt.Printf("%d", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Microsecond)
		fmt.Printf("%c", i)
	}
}

func startMutipleGogroutines() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Microsecond)
	fmt.Println("\ntartMutipleGogroutines")
}

func main() {
	fmt.Println("1. how to start a goroutine")
	startAGoroutine()
	fmt.Println("2. Starting multiple Goroutines")
	startMutipleGogroutines()
}
