package main

import (
	"fmt"
	"math/rand"
	"time"
)

const tamanhoVetor int = 35

var vetor [tamanhoVetor]int

func preencheVetorRandomicamente() {
	var num int
	var numExistente bool
	var numAindaNaoEncontrado bool

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < tamanhoVetor; i++ {
		numAindaNaoEncontrado = true
		for numAindaNaoEncontrado {
			numExistente = false
			num = rand.Intn(tamanhoVetor)
			for j := 0; j < i; j++ {
				if vetor[j] == num {
					numExistente = true
				}
			}
			if !numExistente {
				vetor[i] = num
				numAindaNaoEncontrado = false
			}
		}
	}
}

func insertionSort(a []int, n int) {
	for i := 1; i <= n; i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = key
	}
}

func main() {
	preencheVetorRandomicamente()

	fmt.Println("Vetor original (desordenado): ")
	fmt.Println(vetor)
	fmt.Println()
	s := vetor
	fmt.Println("Vetor ordenado: ")
	insertionSort(s[:], tamanhoVetor-1)
	fmt.Println(s)
}
