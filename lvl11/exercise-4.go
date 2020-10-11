package main

import (
	"fmt"
	"log"
)

type SqrtError struct {
	lat  string
	long string
	err  error
}

func (se SqrtError) Error() string {
	return fmt.Sprintf("Math error-> %v, %v, %v", se.lat, se.long, se.err)
}

func exec4() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
	fmt.Println()
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		numErr := fmt.Errorf("Square root of negative number-> %v", x)
		return 0, SqrtError{"30W", "50E", numErr}
	}
	return 42, nil
}
