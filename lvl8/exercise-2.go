package main

import (
	"encoding/json"
	"fmt"
)

func exec2() {

	type User struct {
		FirstName string `json:"firstName"`
		Age       int    `json:"age"`
	}

	user1 := User{
		FirstName: "Modecai",
		Age:       24,
	}

	user2 := User{
		FirstName: "Rigby",
		Age:       23,
	}

	user3 := User{
		FirstName: "Skips",
		Age:       50,
	}

	users := []User{user1, user2, user3}

	for _, v := range users {

		bytes, _ := json.Marshal(v)
		var reading User
		json.Unmarshal(bytes, &reading)

		fmt.Printf("%+v\n", reading)

	}

	fmt.Println()
}
