package main

import "fmt"

func exec2() {

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for position, value := range slice {
		fmt.Println("My item on", position, "position from my slice has this value-> ", value)
	}

	fmt.Printf("Slice type-> %T\n", slice)

	fmt.Println()
}
