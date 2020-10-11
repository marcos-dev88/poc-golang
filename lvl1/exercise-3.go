package main

import "fmt"

func exec3() {
	var x int = 42
	var y string = "James Bond"
	var z bool = true

	s := fmt.Sprintf("\nx-> %d \ty-> %s \tz-> %t\n", x, y, z)

	fmt.Printf(s)
}
