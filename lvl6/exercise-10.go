package main

import "fmt"

func exec10() {
	firstVar := x()
	secondVar := x()
	fmt.Println("First cloasure-> ", firstVar())
	fmt.Println("First cloasure-> ", firstVar())
	fmt.Println("First cloasure-> ", firstVar())
	fmt.Println("Second cloasure-> ", secondVar())
	fmt.Println("Second cloasure-> ", secondVar())
	fmt.Println("Second cloasure-> ", secondVar())
	fmt.Println()

}

func x() func() string {
	y := 0
	return func() string {
		y++
		str := "This is a cloasure"
		fmt.Print(y, "- ")
		return str

	}
}
