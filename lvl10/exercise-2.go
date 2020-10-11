package main

import "fmt"

func exec2() {

	cs := make(chan int)

	go func() {
		cs <- 42

	}()

	fmt.Println(<-cs)

	fmt.Println("-------")
	fmt.Printf("cs\t%T\n", cs)

	fmt.Println()
}
