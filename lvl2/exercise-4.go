package main

import "fmt"

func exec4() {
	x := 52
	fmt.Printf("Binary-> %b\tDecimal-> %d\tHexadecimal-> %x\n", x, x, x)

	y := x << 1

	fmt.Printf("Binary-> %b\tDecimal-> %d\tHexadecimal-> %x\n", y, y, y)
	fmt.Println()
}
