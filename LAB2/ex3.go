package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "salut"
		ch <- "ce faci?"
		close(ch)
	}()
	for mesaj := range ch {
		fmt.Println(mesaj)
	}
}
