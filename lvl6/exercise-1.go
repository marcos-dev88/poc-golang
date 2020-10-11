package main

import "fmt"

func intExample() int {
	return 10
}

func intStringE() (int, string) {
	return 88, "Miles"
}

func exec1() {

	fmt.Println(intExample())
	fmt.Println(intStringE())

	fmt.Println()
}
