// Primul lucru care se precizeaza este package -ul din care fisierul face
package main

// A doua etapa este importarea pachetelor (nefolosirea unui pachet genereaza
//o eroare)
import "fmt"

// A treia etapa este scrierea functiilor si a tipurilor. Aici avem doar o
//functie main:
func main() {
	var n int
	// var ok int = 0
	ok := 0
	fmt.Scan(&n)
	for i := 2; i < n; i++ {
		if n%i == 0 {
			ok = 1
		}
	}
	if n == 1 {
		ok = 1
	}
	if n == 0 {
		ok = 0
	}
	if ok == 0 {
		fmt.Println("prim")
	} else {
		fmt.Println("nu e prim")
	}
}
