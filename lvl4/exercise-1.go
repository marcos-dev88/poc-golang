package main

import "fmt"

func exec1() {
	arr := [5]int{10, 20, 30, 40, 50}

	for position, value := range arr {
		fmt.Println("My item on", position, "position from my array has this value-> ", value)
	}

	fmt.Printf("Array type-> %T\n", arr)

	fmt.Println()
}
