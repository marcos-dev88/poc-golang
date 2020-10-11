package main

import "fmt"

func exec6() {
	const (
		_ = iota
		b
		c
		d
	)
	fmt.Println(b, c, d)
}
