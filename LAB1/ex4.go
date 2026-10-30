package main

import "fmt"

func main() {
	vector := []float32{9, 11.5, 4.7, 5.12, 2.3}
	var index int
	fmt.Scan(&index)
	if index >= 0 && index < len(vector) {
		vector = append(vector[:index], vector[index+1:]...)
		fmt.Println(vector)
	} else {
		fmt.Println("Index invalid!")
	}
}
