package main

import "fmt"

func main() {
	vector := [5]int{9, 11, 3, 5, 2}
	n := len(vector)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if vector[j] < vector[i] {
				vector[j], vector[i] = vector[i], vector[j]
			}
		}
	}
	fmt.Println(vector)
}
