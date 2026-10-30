package main

import (
	"fmt"
)

type Pozitie struct {
	x int
	y int
}

func main() {
	robot1to2 := make(chan Pozitie)
	robot2to1 := make(chan Pozitie)
	gata := make(chan bool)
	traseu1 := []Pozitie{{0, 0}, {1, 1}, {2, 2}, {3, 3}}
	traseu2 := []Pozitie{{0, 3}, {1, 2}, {2, 1}, {3, 0}}
	go func() {
		for _, pozitie := range traseu1 {
			robot1to2 <- pozitie
		}
		close(robot1to2)
		gata <- true
	}()

	go func() {
		for _, pozitie := range traseu2 {
			robot2to1 <- pozitie
		}
		close(robot2to1)
		gata <- true
	}()
	go func() {
		for pozitie := range robot2to1 {
			fmt.Println("Robot1 a primit:", pozitie)
		}
		gata <- true
	}()
	go func() {
		for pozitie := range robot1to2 {
			fmt.Println("Robot2 a primit:", pozitie)
		}
		gata <- true
	}()
	for i := 0; i < 4; i++ {
		<-gata
	}
	fmt.Println("Ambii roboti au terminat traseul.")
}
