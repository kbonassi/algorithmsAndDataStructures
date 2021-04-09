package main

import "fmt"

const infinito = 4294967295
const null = -1
const qtdVertices = 5

var grafo [qtdVertices][qtdVertices]int

func carregaGrafo() {
	for i := 0; i < qtdVertices; i++ {
		for j := 0; j < qtdVertices; j++ {
			grafo[i][j] = 0
		}
	}

	grafo[0][1] = 1
	grafo[0][3] = 1

	grafo[1][0] = 1
	grafo[1][2] = 1
	grafo[1][3] = 1
	grafo[1][4] = 1

	grafo[2][1] = 1
	grafo[2][4] = 1

	grafo[3][0] = 1
	grafo[3][1] = 1
	grafo[3][4] = 1

	grafo[4][1] = 1
	grafo[4][2] = 1
	grafo[4][3] = 1
}

func enqueue(filaQ []int, elemento int) []int {
	filaQ = append(filaQ, elemento)
	return filaQ
}

func dequeue(fila []int) (int, []int) {
	var filaRetorno = make([]int, 0, 10)
	var elementoRetorno = fila[0]

	for v := range fila {
		if v > 0 {
			filaRetorno = append(filaRetorno, fila[v])
		}
	}
	return elementoRetorno, filaRetorno
}

func bfs(grafo [qtdVertices][qtdVertices]int, verticeInicio int) {
	var verticesCor [qtdVertices]string
	var verticesDistancia [qtdVertices]int
	var verticesPredecessor [qtdVertices]int

	for i := 0; i < qtdVertices; i++ {
		verticesCor[i] = "branco"
		verticesDistancia[i] = infinito
		verticesPredecessor[i] = null
	}

	verticesCor[verticeInicio] = "cinzento"
	verticesDistancia[verticeInicio] = 0
	verticesPredecessor[verticeInicio] = null

	var u int
	var filaQ = make([]int, 0, 10)
	filaQ = enqueue(filaQ, verticeInicio)

	for len(filaQ) != 0 {
		u, filaQ = dequeue(filaQ)
		for v := range grafo[u] {
			if verticesCor[v] == "branco" && grafo[u][v] == 1 {
				verticesCor[v] = "cinzento"
				verticesDistancia[v] = verticesDistancia[u] + 1
				verticesPredecessor[v] = u
				filaQ = enqueue(filaQ, v)
			}
		}
		verticesCor[u] = "preto"
	}

	fmt.Printf("Vértice de partida: %d", verticeInicio)
	fmt.Println()
	fmt.Print("Distância..: ")
	fmt.Println(verticesDistancia)
	fmt.Print("Predecessor: ")
	fmt.Println(verticesPredecessor)
}

func main() {
	carregaGrafo()
	bfs(grafo, 2) // teremos como resultado os vertices alcançados a partir do vértice 2
}
