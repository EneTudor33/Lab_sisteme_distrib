package main

import (
	"fmt"
	"math"
)

type Pozitie struct {
	x float64
	y float64
}

func distanta(a, b Pozitie) float64 {
	return math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

func main() {
	roboti := []Pozitie{
		{0, 0},
		{10, 0},
		{20, 5},
		{20, 20},
		{5, 15},
		{0, 10},
	}
	raza := 20.0
	epsilon := 0.0001
	maxIteratii := 1000
	intalniti := len(roboti) == 1
	for iteratie := 0; iteratie < maxIteratii && !intalniti; iteratie++ {
		pozitiiNoi := make([]Pozitie, len(roboti))
		for i := 0; i < len(roboti); i++ {
			sumaX := roboti[i].x
			sumaY := roboti[i].y
			nr := 1
			for j := 0; j < len(roboti); j++ {
				if i != j && distanta(roboti[i], roboti[j]) < raza {
					sumaX += roboti[j].x
					sumaY += roboti[j].y
					nr++
				}
			}
			nouX := sumaX / float64(nr)
			nouY := sumaY / float64(nr)
			pozitiiNoi[i] = Pozitie{nouX, nouY}
		}
		roboti = pozitiiNoi
		var robotiNoi []Pozitie
		for i := 0; i < len(roboti); i++ {
			exista := false
			for j := 0; j < len(robotiNoi); j++ {
				if distanta(roboti[i], robotiNoi[j]) < epsilon {
					exista = true
				}
			}
			if !exista {
				robotiNoi = append(robotiNoi, roboti[i])
			}
		}
		roboti = robotiNoi
		if len(roboti) == 1 {
			intalniti = true
			fmt.Println("pozitia finala:", roboti[0])
			fmt.Println("numar iteratii:", iteratie+1)
		}
	}
	if !intalniti {
		fmt.Println("nu a fost realizat.")
		fmt.Println("numar roboti activi:", len(roboti))
		fmt.Println("pozitii finale:", roboti)
	}
}
