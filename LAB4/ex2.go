package main

import "fmt"

func GraphFlooding(ids []int, graf [][]int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}
	maxID := make([]int, n)
	for i := 0; i < n; i++ {
		maxID[i] = ids[i]
	}
	for j := 0; j < n-1; j++ {
		urmator := make([]int, n)
		copy(urmator, maxID)
		for i := 0; i < n; i++ {
			for _, vecin := range graf[i] {
				if maxID[i] > urmator[vecin] {
					urmator[vecin] = maxID[i]
				}
			}
		}
		maxID = urmator
	}
	return maxID[0]
}

func main() {
	ids := []int{5, 19, 3, 15}
	graf := [][]int{{1, 2}, {2}, {3}, {0}}

	fmt.Println("Liderul este:", GraphFlooding(ids, graf))
}
