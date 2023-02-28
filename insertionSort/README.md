# InsertionSort

> Exemplo de implementação do algorítmo de ordenação InsertionSort.

## Introdução

O algoritmo InsertionSort e um algoritmo eficiente quando se trata de ordenar um numero pequeno de elementos. Tambem e considerado um algoritmo simples, de facil entendimento.

## Implementação

Neste programa exemplo implementei primeiramente a função _preencheVetorRandomicamente_ para preencher aleatoriamente o array com números entre 0 e o tamanho do array (definido pela constante _tamanhoVetor_), dessa forma, a cada execução do programa será possível notar que o vetor original (desordenado) possui os valores em diferentes índices. Fiz esta função apenas para não criar o vetor com valores fixos.

O algoritmo de ordenaçao em si esta na funçao _insertionSort_.

## Para executar o programa

O algorítmo foi implementado na linguagem _Golang_, sendo portanto esta linguagem necessária para a execução ou compilação da implementação.

Para executar o programa, no shell do seu sistema operacional, estando no diretório fonte, execute:

```shell
$ go run insertionSort.go
```

## Resultado

```shell
Vetor original (desordenado): 
[23 8 2 29 32 11 1 18 21 9 19 17 33 25 26 12 5 3 6 34 14 7 0 24 16 20 27 10 4 22 31 28 13 15 30]

Vetor ordenado: 
[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34]
```
