package main

import "fmt"

func exec5() {
	c := make(chan string)

	go func() {
		c <- "Great scott!"
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)

	fmt.Println()
}
