package main

import "fmt"

func main() {
	persoane := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}
	nume := []string{}
	for persoana := range persoane {
		nume = append(nume, persoana)
	}
	fmt.Println(nume)
	for i := 0; i < len(nume); i++ {
		for j := i + 1; j < len(nume); j++ {
			if persoane[nume[j]] > persoane[nume[i]] {
				nume[j], nume[i] = nume[i], nume[j]
			}
		}
	}
	fmt.Println(nume)
}
