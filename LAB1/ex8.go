package main

import (
	"fmt"
)

func main() {
	model := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}
	fmt.Println(model)
	raspuns := make(map[int][]string)
	for i, j := range model {
		raspuns[j] = append(raspuns[j], i)
	}
	fmt.Println(raspuns)
}
