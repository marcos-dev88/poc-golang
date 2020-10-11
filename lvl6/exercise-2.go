package main

import "fmt"

func variadicInt(x ...int) int {
	n := 0

	for _, v := range x {
		n += v
	}
	return n
}

func sliceOfInt(x []int) int {
	n := 0

	for _, v := range x {
		n += v
	}
	return n
}

func exec2() {
	s := []int{2, 3, 4, 5, 7, 3, 4}
	resultFirst := variadicInt(s...)
	fmt.Println(resultFirst)

	slic := []int{12, 34, 2, 1, 2}
	fmt.Println(sliceOfInt(slic))

	fmt.Println()
}
