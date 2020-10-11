package main

import (
	"fmt"
	"runtime"
)

var count int = 0

func exec3() {
	totalGoRoutines := 10
	goRoutines(totalGoRoutines)
	wg.Wait()

	fmt.Println("Go Routines total-> ", totalGoRoutines, "\nCount total->", count)
	fmt.Println()
}

func goRoutines(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func() {
			goRoutinesCount := count
			runtime.Gosched()
			goRoutinesCount++
			count = goRoutinesCount
			wg.Done()
		}()
	}
}
