package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func exec1() {

	type Person struct {
		FirstName string
		LastName  string
		Sayings   []string
	}

	p1 := Person{
		FirstName: "James",
		LastName:  "Bond",
		Sayings:   []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))

	fmt.Println()
}
