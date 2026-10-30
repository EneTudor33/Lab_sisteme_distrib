package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Eroare la deschiderea fisierului:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	cuvinte := []string{}

	for scanner.Scan() {
		cuvinte = append(cuvinte, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Eroare la citire:", err)
		return
	}
	fmt.Println(cuvinte)
	frecventa := make(map[string]int)
	for _, cuvant := range cuvinte {
		frecventa[cuvant]++
	}
	fmt.Println(frecventa)
}
