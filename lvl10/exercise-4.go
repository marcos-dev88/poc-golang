package main

import "fmt"

func exec4() {
	a := make(chan int)
	b := make(chan int)
	x := 10

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("This is a value from channel A-> ", v)
		case v := <-b:
			fmt.Println("This is a value from channel B-> ", v)
		}
	}
	fmt.Println()
}
