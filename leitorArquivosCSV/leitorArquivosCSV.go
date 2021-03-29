package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Estado é uma estrutura para receber os registros do arquivo CSV
type Estado struct {
	estado  string
	capital string
	sigla   string
}

// checkErr é uma função para tratamento de erro
func checkErr(err error) {
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}

// CSV é uma função para ler o arquivo CSV
func CSV(fileName string) []Estado {
	var vestado string
	var vcapital string
	var vsigla string
	estados := make([]Estado, 0, 10)

	file, err := os.Open(fileName)
	if err != nil {
		log.Panic("ERRO: ", err.Error())
	}

	reader := csv.NewReader(bufio.NewReader(file))
	// define o delimitador dos registros CSV como sendo ','
	reader.Comma = ','

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			checkErr(err)
		}

		vestado = record[0]
		vcapital = record[1]
		vsigla = record[2]
		
		// despreza o header do arquivo CSV
		if string(vestado[0]) == "#" {
			continue
		}
		
		estados = append(estados, Estado{
			estado:  vestado,
			capital: vcapital,
			sigla:   vsigla,
		})
	}

	return estados
}

func main() {
	slice := CSV("Estados.csv")

	fmt.Println(slice)
}

