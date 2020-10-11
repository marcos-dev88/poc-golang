package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) toTalk() {
	fmt.Println("Hello my name is", p.name, "and i have", p.age, "years old")
}

type humans interface {
	toTalk()
}

func saySomething(h humans) {
	h.toTalk()
}

func exec2() {

	someone := Person{
		name: "Mordecai",
		age:  23,
	}

	saySomething(&someone)

	fmt.Println()

}
