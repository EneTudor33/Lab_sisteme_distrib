package main

import "fmt"

func AverageFlooding(voturi []float64, graf [][]int) float64 {
	n := len(voturi)
	if n == 0 {
		return 0
	}
	valori := make([]float64, n)
	copy(valori, voturi)
	for j := 0; j < 1000; j++ {
		urmator := make([]float64, n)
		for i := 0; i < n; i++ {
			suma := valori[i]
			for _, vecin := range graf[i] {
				suma += valori[vecin]
			}
			urmator[i] = suma / float64(len(graf[i])+1)
		}
		valori = urmator
		minim := valori[0]
		maxim := valori[0]
		for i := 1; i < n; i++ {
			if valori[i] < minim {
				minim = valori[i]
			}
			if valori[i] > maxim {
				maxim = valori[i]
			}
		}
		if maxim-minim < 0.0001 {
			break
		}
	}
	return valori[0]
}

func main() {
	voturi := []float64{1, 1, 0, 1}
	graf := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	rezultat := AverageFlooding(voturi, graf)
	fmt.Println("Media:", rezultat)
	if rezultat > 0.5 {
		fmt.Println("Majoritatea a votat DA")
	} else if rezultat < 0.5 {
		fmt.Println("Majoritatea a votat NU")
	} else {
		fmt.Println("Egalitate")
	}
}
