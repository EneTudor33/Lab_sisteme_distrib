package main

import "fmt"

func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	fmt.Println(slice)
}
