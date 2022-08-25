package main

import (
	"fmt"
	"os"
	"strconv"
)

type People struct {
	name    string
	address string
	job     string
	reason  string
}

func main() {
	test := os.Args
	index, _ := strconv.Atoi(test[1])
	persons := []*People{
		{name: "Inas", address: "a", job: "a", reason: "xxx"},
		{name: "Ricky", address: "b", job: "a", reason: "xxx"},
		{name: "Thalia", address: "c", job: "a", reason: "xxx"},
		{name: "Yudha", address: "d", job: "a", reason: "xxx"},
		{name: "Clara", address: "e", job: "a", reason: "xxx"},
		{name: "Agus", address: "f", job: "a", reason: "xxx"},
		{name: "Yosef", address: "g", job: "a", reason: "xxx"},
		{name: "Eka", address: "h", job: "a", reason: "xxx"},
		{name: "Satrio", address: "i", job: "a", reason: "xxx"},
		{name: "Rijal", address: "j", job: "a", reason: "xxx"},
	}
	// getData()
	fmt.Println(persons[index].name, persons[index].address, persons[index].job, persons[index].name, persons[index].reason)
}
