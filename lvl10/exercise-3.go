package main

import "fmt"

func exec3() {
	ch := genericFunc()
	recieve(ch)

	fmt.Println()
}

func genericFunc() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func recieve(r <-chan int) {
	for v := range r {
		fmt.Println(v)
	}

}
