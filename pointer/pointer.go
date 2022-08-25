package main

import "fmt"

type People struct {
	name   string
	gender string
}

func main() {
	persons := []*People{
		{name: "Inas", gender: "P"},
		{name: "Ricky", gender: "L"},
		{name: "Thalia", gender: "P"},
		{name: "Yudha", gender: "L"},
		{name: "Clara", gender: "P"},
		{name: "Agus", gender: "L"},
		{name: "Yosef", gender: "L"},
		{name: "Eka", gender: "L"},
		{name: "Satrio", gender: "L"},
		{name: "Rijal", gender: "L"},
	}

	printFriends := func(friends []*People) {
		for _, value := range friends {
			fmt.Printf("Name: %v Gender: %v \n", value.name, value.gender)
		}
	}

	printFriends(persons)
}
