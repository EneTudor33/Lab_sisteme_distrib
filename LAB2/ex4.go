package main

import "fmt"

func main() {
	gata := make(chan bool)
	go func() {
		for i := 1; i <= 67; i++ {
			fmt.Println(i)
		}
		gata <- true
	}()
	var m bool
	m = <-gata
	fmt.Println("gata", m)
}
