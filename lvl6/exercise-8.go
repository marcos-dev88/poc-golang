package main

import "fmt"

func funcExample() func() {
	return func() {
		fmt.Println("We got it! :D")
	}

}

func exec8() {
	x := funcExample()

	x()

	fmt.Println()
}
