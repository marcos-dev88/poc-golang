package main

import (
	"fmt"
)

func exec7() {
	rangeChan := 100
	c := make(chan int)
	goRoutines(10, c)

	for i := 0; i < rangeChan; i++ {
		fmt.Println(i, "\t", <-c)
	}

	fmt.Println()
}

func goRoutines(num int, ch chan int) {
	for i := 0; i < num; i++ {
		go func() {
			for j := 0; j < num; j++ {
				ch <- j
			}
		}()
	}
}
