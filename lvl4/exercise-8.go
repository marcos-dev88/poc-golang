package main

import "fmt"

func exec8() {
	myTable := map[string][]string{
		"brown_helmet": []string{
			"He likes to create new thing", "Build the time machine",
		},
		"mcfly_marty": []string{
			"He likes to play guitar", "Help the Doc with his invensions",
		},
		"tannen_biff": []string{
			"Make bully with Geaorge McFly", "Seems like he loves manure",
		},
	}

	for key, v := range myTable {
		fmt.Println(key)
		for _, value := range v {
			fmt.Println("\t", " - ", value)
		}
	}

	fmt.Println("\n", myTable)

	fmt.Println()
}
