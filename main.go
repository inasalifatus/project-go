package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// Tugas 1 : Looping
	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("Ganjil", i)
		} else {
			fmt.Println("Genap", i)
		}
	}
}
