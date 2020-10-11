package main

import "fmt"

func exec1() {

	type person struct {
		name        string
		lastName    string
		favIceCream []string
	}

	benson := person{
		name:        "Benson",
		lastName:    "Dunwoody",
		favIceCream: []string{"Vanilla", "Chocolate"},
	}

	fmt.Println(benson.name, benson.lastName)
	for p, _ := range benson.favIceCream {
		fmt.Println("\t-", benson.favIceCream[p])

	}
	fmt.Println()

	walks := person{
		name:        "Walks",
		lastName:    "Quippenger",
		favIceCream: []string{"Vanilla", "Banana"},
	}

	fmt.Println(walks.name, walks.lastName)
	for p, _ := range walks.favIceCream {
		fmt.Println("\t-", walks.favIceCream[p])
	}

	fmt.Println()
}
