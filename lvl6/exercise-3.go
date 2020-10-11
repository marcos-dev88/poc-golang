package main

import "fmt"

func showLineb4() string {
	return "It's heavy!"
}

func showLineAfter() string {
	return "Great Scott!"
}

func exec3() {

	defer fmt.Println(showLineAfter(), "\n-")

	fmt.Println(showLineb4())

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()
}
