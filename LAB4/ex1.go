package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool // liderul anunta ca a fost ales
}

// realizeaza operatiile interne nodului
func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- Msg{Value: id, Done: false}
	for {
		msg := <-in
		if msg.Done {
			if *isLeader {
				return
			}
			out <- msg
			return
		}
		if msg.Value > id {
			out <- msg
		} else if msg.Value == id {
			*isLeader = true
			out <- Msg{Value: id, Done: true}
		}
	}
}

// returneaza id -ul liderului
func RingFlooding(ids []int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}
	channels := make([]chan Msg, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan Msg, n+1)
	}
	leaders := make([]bool, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go node(ids[i], channels[i], channels[(i+1)%n], &leaders[i], &wg)
	}
	wg.Wait()
	for i := 0; i < n; i++ {
		if leaders[i] {
			return ids[i]
		}
	}
	return -1
}

func main() {
	ids := []int{5, 19, 3, 15}
	fmt.Println("Liderul este:", RingFlooding(ids))
}
