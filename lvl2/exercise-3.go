package main

import "fmt"

func exec3() {
	const a = 10
	const b int = 12

	fmt.Printf("\na type-> %T, a value-> %v\t b type-> %T, b value -> %v\n", a, a, b, b)
	fmt.Println()
}
