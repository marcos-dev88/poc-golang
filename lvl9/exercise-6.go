package main

import (
	"fmt"
	"runtime"
)

func exec6() {
	fmt.Println("Your OS is-> ", runtime.GOOS)
	fmt.Println("Your ARCH is-> ", runtime.GOARCH)
	fmt.Println()
}
