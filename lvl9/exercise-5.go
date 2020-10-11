package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

var count3 int64

func exec5() {
	totalGoRoutines := 10
	wg.Add(totalGoRoutines)

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&count3, 1)
			runtime.Gosched()
			fmt.Println("Count-> ", atomic.LoadInt64(&count3))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println()
}
