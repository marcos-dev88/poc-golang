package main

import "fmt"

func exec5() {
	x := 10
	y := 100

	for x <= y {
		z := x % 4

		fmt.Println(z)
		x++
	}
	fmt.Println()
}
