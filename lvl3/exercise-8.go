package main

import "fmt"

func exec8() {

	//pPoints1 := 200
	//pPoints2 := 100

	//pPoints1 := 100
	//pPoints2 := 200

	pPoints1 := 100
	pPoints2 := 100

	switch {
	case pPoints1 > pPoints2:
		fmt.Println("Player 1 was the Winner!")
	case pPoints1 < pPoints2:
		fmt.Println("Player 2 was the Winner!")
	default:
		fmt.Println("Tie!")
	}

	fmt.Println()
}
