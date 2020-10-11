package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func exec1() {
	wg.Add(2)

	go func() {
		for i := 0; i <= 20; i++ {
			fmt.Println("First Go Routine->", i)
			time.Sleep(100)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= 20; i++ {
			fmt.Println("Second Go Routine->", i)
			time.Sleep(100)
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println()
}
