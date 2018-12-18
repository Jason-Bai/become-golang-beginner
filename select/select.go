package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func selectExample() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

func process(ch chan string) {
	time.Sleep(10500 * time.Microsecond)
	ch <- "process successful"
}

func defaultSelect() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Microsecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}

func server11(ch chan string) {
	ch <- "from server11"
}

func server22(ch chan string) {
	ch <- "from server22"
}

func randomSelect() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server11(output1)
	go server22(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

func main() {
	fmt.Println("1. select example")
	selectExample()
	fmt.Println("2. default select")
	defaultSelect()
	fmt.Println("3. random select")
	randomSelect()
}
