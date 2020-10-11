package main

import "fmt"

func exec6() {
	func() {
		fmt.Println("Hello Anonymous func :D")
	}()

	fmt.Println()
}
