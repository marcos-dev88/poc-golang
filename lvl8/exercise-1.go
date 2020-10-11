package main

import (
	"encoding/json"
	"fmt"
)

func exec1() {

	type User struct {
		FirstName string `json:"firstName"`
		Age       int    `json:"age"`
	}

	user1 := User{
		FirstName: "James",
		Age:       32,
	}

	user2 := User{
		FirstName: "Moneypenny",
		Age:       27,
	}

	user3 := User{
		FirstName: "M",
		Age:       54,
	}

	users := []User{user1, user2, user3}

	fmt.Println(users)

	bytes, err := json.MarshalIndent(users, "\t", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))

	fmt.Println()
}
