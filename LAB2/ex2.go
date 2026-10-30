package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		mesaj := <-ch
		fmt.Println(mesaj)
	}()
	ch <- "salut"
}

//cu mentiune ca din toate incercarile de a rula codul mereu main s-a terminat inainte de afisare
