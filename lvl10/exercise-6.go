package main

import "fmt"

func exec6() {
	c := make(chan int)
	go sendChanel(20, c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println()
}

func sendChanel(num int, ch chan int) {
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
}
