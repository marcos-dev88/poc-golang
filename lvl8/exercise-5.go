package main

import (
	"fmt"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	Sayings   []string
}

type orderByLastName []User

func (ln orderByLastName) Len() int {
	return len(ln)
}

func (ln orderByLastName) Less(i, j int) bool {
	return ln[i].FirstName < ln[j].FirstName
}

func (ln orderByLastName) Swap(i, j int) {
	ln[i], ln[j] = ln[j], ln[i]
}

// ========================================

type orderByAge []User

func (a orderByAge) Len() int {
	return len(a)
}

func (a orderByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a orderByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func exec5() {

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
		LastName:  "Hmmmm",
		Age:       54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(orderByLastName(users))
	fmt.Println("Sorted by LastName: ")
	for p, v := range users {
		fmt.Println(" ", p+1, "User: \n", "   Full Name:", v.FirstName, v.LastName, "| Age:", v.Age)
		fmt.Println("      Sayings:")
		for _, v := range users[p].Sayings {
			fmt.Println("\t-", v)
		}
	}

	fmt.Println()

	sort.Sort(orderByAge(users))
	fmt.Println("Sorted by Age: ")
	for p, v := range users {
		fmt.Println(" ", p+1, "User: \n", "   Full Name:", v.FirstName, v.LastName, "| Age:", v.Age)
		fmt.Println("      Sayings:")
		for _, v := range users[p].Sayings {
			fmt.Println("\t-", v)
		}
	}

	fmt.Println()
}
