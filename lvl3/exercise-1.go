package main

import "fmt"

func exec1() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
	fmt.Println()
}
