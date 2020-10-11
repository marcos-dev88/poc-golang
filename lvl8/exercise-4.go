package main

import (
	"fmt"
	"sort"
)

func exec4() {
	intSlice := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}

	stringSlice := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println("Slice of int not sorted-> ", intSlice)

	fmt.Println("Slice of string not sorted-> ", stringSlice)

	sort.Ints(intSlice)
	fmt.Println("Slice of int sorted-> ", intSlice)

	sort.Strings(stringSlice)
	fmt.Println("Slice of strinf sorted-> ", stringSlice)

	fmt.Println()
}
