package main

import "fmt"

func exec7() {

	//kidAge := 10
	kidAge := 16
	//kidAge := 18

	if kidAge < 13 {
		fmt.Println("Sorry kid buy you can't enter on roller coaster")
	} else if kidAge >= 13 && kidAge < 17 {
		fmt.Println("Okay, you can pass, but be careful and hold up hard on seat")
	} else {
		fmt.Println("Okay kid, you can pass. Enjoy!")
	}

	fmt.Println()
}
