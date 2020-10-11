package main

import "fmt"

type Person struct {
	name     string
	lastName string
	age      int
}

func (p Person) showNameAndAge() {
	fmt.Println("Hello my name is", p.name, p.lastName, "and my age is", p.age)
}

func exec4() {
	adam := Person{
		name:     "Adam",
		lastName: "Evans",
		age:      16,
	}

	adam.showNameAndAge()
	fmt.Println()
}
