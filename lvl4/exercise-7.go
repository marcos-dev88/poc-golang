package main

import "fmt"

func exec7() {
	myTable := [][]string{
		[]string{
			"Helmet",
			"Brown",
			"likes to create new things and builded the time machine",
		},
		[]string{
			"Marty",
			"McFly",
			"likes to play guitar and help the Doc with his invensions",
		},
		[]string{
			"Biff",
			"Tannen",
			"seems like he loves manure",
		},
	}

	for _, v := range myTable {
		fmt.Println(v)
	}

	fmt.Println()

}
