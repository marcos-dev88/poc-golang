package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func exec3() {
	type User struct {
		FirstName string
		LastName  string
		Age       int
		Sayings   []string
	}

	u1 := User{
		FirstName: "James",
		LastName:  "Bond",
		Age:       32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		FirstName: "Miss",
		LastName:  "Moneypenny",
		Age:       27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		FirstName: "M",
		LastName:  "Hmmm",
		Age:       54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}
	//fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println()
}
