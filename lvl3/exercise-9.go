package main

import "fmt"

func exec9() {
	esporteFavorito := "volleyball"
	sentence := "Your favorite sport is"

	switch esporteFavorito {
	case "basketball":
		fmt.Println(sentence, "Basketball")
	case "football":
		fmt.Println(sentence, "Football")
	case "chess":
		fmt.Println(sentence, "Chess")
	case "volleyball":
		fmt.Println(sentence, "Voleyball")
	case "skate":
		fmt.Println(sentence, "Skate")
	}

}
