package main

import (
	"fmt"
	"time"
)

func main() {
	robot1 := make(chan bool)
	robot2 := make(chan bool)
	go func() {
		time.Sleep(6 * time.Second)
		robot1 <- true
	}()

	go func() {
		time.Sleep(3 * time.Second)
		robot2 <- true
	}()
	select {
	case <-robot1:
		fmt.Println("Robot1")
	case <-robot2:
		fmt.Println("Robot2")
	}
}
