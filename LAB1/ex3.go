package main

import "fmt"

func main() {
	vector := [5]int{9, 11, 4, 5, 2}
	n := len(vector)
	pare := []int{}
	for i := 0; i < n; i++ {
		if vector[i]%2 == 0 {
			pare = append(pare, vector[i])
		}
	}
	fmt.Println(pare)
}
