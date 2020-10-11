package main

import "fmt"

type Person struct {
	name     string
	lastName string
}

func changeMe(p *Person) {
	p.name = "Eileen"
	p.lastName = "Roberts"
	fmt.Println("I have changed that person for this:", p.name, p.lastName)
}

func exec2() {
	personEx := Person{
		name:     "Someone",
		lastName: "SomeLastName",
	}

	changeMe(&personEx)

	fmt.Println()
}
