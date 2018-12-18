package main

import (
	"fmt"
	"time"
)

func declaringChannels() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
}

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func blockingByDefault() {
	done := make(chan bool)
	fmt.Println("call hello go goroutine")
	go hello(done)
	<-done
	fmt.Println("blockingByDefault received data")
}

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func blockingMutipleChannels() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func closeChannel() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received: ", v, ok)
	}
}

func closeChannelByRange() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received: ", v)
	}
}

func main() {
	fmt.Println("1. declaring channels")
	declaringChannels()
	fmt.Println("2.Sends and receives are blocking by default ")
	blockingByDefault()
	fmt.Println("3. another example for channels")
	blockingMutipleChannels()
	fmt.Println("4. Closing channels and for range loops on channels")
	closeChannel()
	closeChannelByRange()
}
