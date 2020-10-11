package main

import "fmt"

type SpecialError struct {
	someError string
}

func (e SpecialError) Error() string {
	return fmt.Sprintf("We found some error around here-> %v", e.someError)
}

func parameterError(err error) {
	fmt.Println("Someone send this error with a parameter-> ", err)
}

func exec3() {
	arg := SpecialError{"##>#<"}
	parameterError(arg)
	fmt.Println()
}
