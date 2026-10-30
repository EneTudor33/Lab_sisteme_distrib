package main

import "fmt"

func send(ch chan string) {
	ch <- "send"
}
func receive(ch chan string) string {
	mesaj := <-ch
	return mesaj
}
func main() {
	ch := make(chan string)
	rezultat := make(chan string)
	go send(ch)
	go func() {
		mesaj := receive(ch)
		rezultat <- mesaj
	}()
	rezultat2 := <-rezultat
	fmt.Println(rezultat2)
}
