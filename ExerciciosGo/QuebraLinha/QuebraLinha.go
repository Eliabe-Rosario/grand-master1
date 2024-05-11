package main

import (
	"fmt"
	"strings"
)

func quebrarLinha(frase string, colunas int) string {
	palavras := strings.Fields(frase)
	var resultado string
	comprimentoLinha := 0

	for _, palavra := range palavras {
		if comprimentoLinha+len(palavra) <= colunas {
			resultado += palavra + " "
			comprimentoLinha += len(palavra) + 1
		} else {
			resultado += "\n" + palavra + " "
			comprimentoLinha = len(palavra) + 1
		}
	}

	return strings.TrimSpace(resultado)
}

func main() {
	var frase string
	var colunas int

	fmt.Print("Digite a frase: ")
	fmt.Scanln(&frase)

	fmt.Print("Digite a quantidade de colunas: ")
	fmt.Scanln(&colunas)

	fraseFormatada := quebrarLinha(frase, colunas)
	fmt.Println(fraseFormatada)
}
