package main

import "fmt"

func showCallBack() {
	fmt.Println("it's a func of callback")
	fmt.Println()
}

func countAndCallBack(f func()) {
	fmt.Println("Look at this:")
	f()
}

func exec9() {
	countAndCallBack(showCallBack)
}
