package main

import "fmt"

func exec1() {

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	fmt.Println()
}
