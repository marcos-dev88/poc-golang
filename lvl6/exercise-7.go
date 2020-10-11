package main

import "fmt"

func exec7() {

	variableFunc := func() string {
		return "Some variable with function :b"
	}()

	fmt.Println(variableFunc)

	fmt.Println()
}
