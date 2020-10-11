package main

import "fmt"

func exec2() {

	type person struct {
		name        string
		lastName    string
		favIceCream []string
	}

	myMap := make(map[string]person)

	myMap["fives"] = person{
		name:        "High",
		lastName:    "Fives",
		favIceCream: []string{"Mint", "Pineaple"},
	}

	myMap["mitch"] = person{
		name:        "Mitch",
		lastName:    "Sorestein",
		favIceCream: []string{"Chocolate", "Mint"},
	}

	for _, v := range myMap {
		fmt.Println(v.name, v.lastName)
		for p, _ := range v.favIceCream {
			fmt.Println("\t-", v.favIceCream[p])
		}
		fmt.Println()
	}

	fmt.Println()
}
