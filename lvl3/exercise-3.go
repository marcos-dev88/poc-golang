package main

import "fmt"

func exec3() {
	birthYear := 1999
	actualYear := 2020

	for birthYear <= actualYear {
		fmt.Println(birthYear)
		birthYear++
	}
	fmt.Println()
}
