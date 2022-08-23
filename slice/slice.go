package main

import "fmt"

func main() {
	classmates := []string{"inas", "ricky", "thalia", "yudha", "clara", "agus", "yosef", "eka", "satrio", "rijal"}
	for index, element := range classmates {
		fmt.Println(index, element)
	}

}
