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

func quickSort(p int, r int) {
	if p < r {
		q := partition(p, r)
		quickSort(p, q-1)
		quickSort(q+1, r)
	}
}

func partition(p int, r int) int {
	x := vetor[r]
	i := p - 1
	var troca int
	for j := p; j <= r-1; j++ {
		if vetor[j] <= x {
			i = i + 1
			troca = vetor[i]
			vetor[i] = vetor[j]
			vetor[j] = troca
		}
	}
	troca = vetor[i+1]
	vetor[i+1] = vetor[r]
	vetor[r] = troca
	return i + 1
}

func main() {
	preencheVetorRandomicamente()

	fmt.Println("Vetor original (desordenado): ")
	fmt.Println(vetor)
	fmt.Println()
	fmt.Println("Vetor ordenado: ")
	quickSort(0, tamanhoVetor-1)
	fmt.Println(vetor)
}
