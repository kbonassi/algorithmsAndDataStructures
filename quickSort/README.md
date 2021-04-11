# QuickSort

> Exemplo de implementação do algorítmo de ordenação QuickSort.

## Introdução

QuickSort é um dos algorítmos de ordenação mais populares e eficazes, sendo amplamente ensinado nos cursos de algorítmos e estrutura de dados, bem como utilizado em muitas situações.

O objetivo do QuickSort é ordenar dados utilizando-se do paradigma de divisão e conquista,
onde o arranjo é quebrado em pedaços para realizar a ordenação.

## Implementação

Neste programa exemplo implementei primeiramente a função _preencheVetorRandomicamente_ para preencher aleatoriamente o array com números entre 0 e o tamanho do array (definido pela constante _tamanhoVetor_), dessa forma, a cada execução do programa será possível notar que o vetor original (desordenado) possui os valores em diferentes índices. Fiz esta função apenas para não colocar o vetor com valores fixos.

O algorítmo em si é sustentado pelas funções _quickSort_ e _partition_ as quais realizam a ordenação.

## Para executar o programa

O algorítmo foi implementado na linguagem _Golang_, sendo portanto esta linguagem necessária para a execução ou compilação da implementação.

Para executar o programa, no shell do seu sistema operacional, estando no diretório fonte, comande:

```shell
$ go run quickSort.go
```

## Resultado

```shell
Vetor original (desordenado): 
[1 23 33 21 27 30 20 22 19 4 28 17 3 5 6 0 8 14 31 29 9 11 24 34 18 2 10 13 26 7 25 32 15 16 12]

Vetor ordenado: 
[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34]
```
