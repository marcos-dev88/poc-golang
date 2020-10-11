package main

import "fmt"

func exec4() {

	food := struct {
		name        string
		isSweet     bool
		placesToBuy []string
		goodToEat   map[string]string
	}{
		name:    "Pizza",
		isSweet: false,
		placesToBuy: []string{
			"Pizza Parlor",
			"Pizza Restaurant",
			"Pizza Placê",
			"Pizza Shop",
		},
		goodToEat: map[string]string{
			"breakfast": "Yes, there's people who eats pizza on breakfast, but not everyday",
			"lunch":     "Yes! It's a complete meal, but not eat much, because it's a fast food. It's better to eat a salda with beef",
			"snack":     "No, because it's greasy... It's better to eat in important meals.",
			"dinner":    "Yes! It's perfect to eat in dinner!",
		},
	}

	fmt.Println(food.name)
	if food.isSweet == true {
		fmt.Println("Sweet Pizza ")
	} else {
		fmt.Println("Savory Pizza")
	}
	fmt.Println("We could buy this food in:")
	for _, v := range food.placesToBuy {
		fmt.Println("\t-", v)
	}
	fmt.Println("This kind of food is good to eat on:")
	for key, value := range food.goodToEat {
		fmt.Println("\t", key, " - ", value)
	}

	fmt.Println()

}
