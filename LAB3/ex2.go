package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type LamportClock struct{ counter int }

// incrementeaza contorul
func (l *LamportClock) Tick() {
	l.counter++
}
func (l *LamportClock) Send() {
	l.Tick()
}
func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.Tick()
}

type Message struct{ sender, ts int }
type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

// TODO: retuneaza true daca a precede b
func less(a, b Event) bool { return a.internal_counter < b.internal_counter }

func main() {
	events := make(chan Event, 6)
	ch12 := make(chan Message)
	ch23 := make(chan Message)
	ch31 := make(chan Message)
	gata := make(chan bool)
	go func() {
		clock := LamportClock{}
		n := rand.Intn(4)
		for i := 0; i < n; i++ {
			clock.Tick()
		}
		clock.Send()
		ch12 <- Message{1, clock.counter}
		events <- Event{1, clock.counter, "send"}
		n = rand.Intn(4)
		for i := 0; i < n; i++ {
			clock.Tick()
		}
		msg := <-ch31
		clock.Receive(msg.ts)
		events <- Event{1, clock.counter, "recv"}
		gata <- true
	}()
	go func() {
		clock := LamportClock{}
		n := rand.Intn(4)
		for i := 0; i < n; i++ {
			clock.Tick()
		}
		msg := <-ch12
		clock.Receive(msg.ts)
		events <- Event{2, clock.counter, "recv"}
		n = rand.Intn(4)
		for i := 0; i < n; i++ {
			clock.Tick()
		}
		clock.Send()
		events <- Event{2, clock.counter, "send"}
		ch23 <- Message{2, clock.counter}
		gata <- true
	}()
	go func() {
		clock := LamportClock{}
		n := rand.Intn(4)
		for i := 0; i < n; i++ {
			clock.Tick()
		}
		msg := <-ch23
		clock.Receive(msg.ts)
		events <- Event{3, clock.counter, "recv"}
		n = rand.Intn(4)
		for i := 0; i < n; i++ {
			clock.Tick()
		}
		clock.Send()
		events <- Event{3, clock.counter, "send"}
		ch31 <- Message{3, clock.counter}
		gata <- true
	}()
	for i := 0; i < 3; i++ {
		<-gata
	}
	var listaEvenimente []Event
	for i := 0; i < 6; i++ {
		eveniment := <-events
		listaEvenimente = append(listaEvenimente, eveniment)
	}
	sort.Slice(listaEvenimente, func(i, j int) bool {
		return less(listaEvenimente[i], listaEvenimente[j])
	})
	for _, e := range listaEvenimente {
		fmt.Println("Proces:", e.process, "Timestamp:", e.internal_counter, "Tip:", e.kind)
	}
}
