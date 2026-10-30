package main

import (
	"fmt"
)

func main() {
	input := []int{2, 7, 11, 8, 4, 9}
	var s int
	fmt.Scan(&s)
	verif := make(map[int]int)
	for i, num := range input {
		complement := s - num
		index, ok := verif[complement]
		if ok {
			fmt.Println(i, index)
		}
		verif[num] = i
	}
}
