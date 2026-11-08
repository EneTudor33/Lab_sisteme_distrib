package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

// TODO:
// Calculeaza decalajul theta si round -trip -ul delta
func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	delta = t4.Sub(t1) - t3.Sub(t2)
	return theta, delta
}
func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()
	// Procesare intre 5 si 20 ms
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()
	// Response
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()
	return estimate(t1, t2, t3, t4)
}
func main() {
	client := Clock{offset: -100 * time.Millisecond}
	server := Clock{offset: 0 * time.Millisecond}
	fmt.Println("Offset initial:", client.offset)
	theta, delta := measure(&client, &server)
	fmt.Println("Theta:", theta)
	fmt.Println("Delta:", delta)
	client.offset += theta
	fmt.Println("Ceasul ajustat", client.offset)
	thetaInitial := theta
	deltaInitial := delta
	var thetaMin time.Duration
	var deltaMin time.Duration
	for i := 0; i < 10; i++ {
		client.offset = -100 * time.Millisecond
		fmt.Println("Offset initial:", client.offset)
		theta, delta := measure(&client, &server)
		fmt.Println("Theta:", theta)
		fmt.Println("Delta:", delta)
		if i == 0 || delta < deltaMin {
			deltaMin = delta
		}
		if i == 0 || theta < thetaMin {
			thetaMin = theta
		}
	}
	fmt.Println("Theta minim:", thetaMin)
	fmt.Println("Theta initial:", thetaInitial)
	fmt.Println("Delta minim:", deltaMin)
	fmt.Println("Delta initial:", deltaInitial)
}
