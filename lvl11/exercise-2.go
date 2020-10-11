package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func exec2() {
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

	bs, err := toJson(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
	fmt.Println()
}

func toJson(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("We've some problem to convert to json")
	}
	return bs, nil
}
