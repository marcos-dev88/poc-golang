package main

import "fmt"

type steak int //Recieving a int value to type 'steak'

func exec4() {
	var x steak
	fmt.Printf("\nx value-> %v \t x type-> %T \n", x, x)
	x = 42
	fmt.Println("Putting value '42' to x->", x)
	fmt.Println()
}
