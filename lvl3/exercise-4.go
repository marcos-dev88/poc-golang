package main

import "fmt"

func exec4() {
	actualYear := 2020
	birthDate := 1999

	for {
		if birthDate > actualYear {
			break
		}
		fmt.Println(birthDate)
		birthDate++
	}
	fmt.Println()
}
