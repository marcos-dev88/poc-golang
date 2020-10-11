package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mutex sync.Mutex
var count2 int

func exec4() {

	totalGoRoutines := 10
	goRoutines2(totalGoRoutines)
	wg.Wait()

	fmt.Println("Go Routines total-> ", totalGoRoutines, "\nCount total->", count2)
	fmt.Println()

}

func goRoutines2(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func() {
			mutex.Lock()
			goRoutinesCount := count2
			runtime.Gosched()
			goRoutinesCount++
			count2 = goRoutinesCount
			mutex.Unlock()
			wg.Done()
		}()
	}
}
