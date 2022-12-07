package main

import (
	arq "BuscaMusica/leituraescrita"
	mapa "BuscaMusica/mapa"
	ord "BuscaMusica/ordenacao"
	hash "BuscaMusica/tabelahash"
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	dados := arq.LeituraJson(`arquivos/songs7JSONvector.txt`)

	dados = ord.CountingSort(dados)

	mapa.CriaMapaArquivos(dados)
	tabela := hash.TabelaHash{}
	tabela.FormaTabelaHash(dados)

	fmt.Print("\n----- BUSCA DE NOTAS EM ARQUIVO DE MÚSICA -----")

	for {

		var arquivo string
		fmt.Print("\n\nQual o número do arquivo que deseja consultar? ")
		fmt.Scan(&arquivo)

		if regexp.MustCompile("[0-9]+").FindString(arquivo) != "" {

			numArq, _ := strconv.Atoi(arquivo)

			var nota string
			fmt.Print("\nQual a nota que deseja pesquisar? ")
			fmt.Scan(&nota)

			if regexp.MustCompile("[A-G]").FindString(nota) != "" {

				var opcao string
				fmt.Print("\nBuscar no arquivo ", arquivo, ":")
				fmt.Print("\n\n1 - Quantas notas ", nota, " existem")
				fmt.Print("\n2 - Se contém a nota ", nota, " em uma determinada linha")
				fmt.Print("\n\nInsira a opção que deseja: ")
				fmt.Scan(&opcao)

				if opcao == "1" {

					quantidade, _ := tabela.BuscarNotaArquivo(numArq, nota)
					// quantidade, linhas := tabela.BuscarNotaArquivo(numArq, nota)

					if quantidade > 0 {

						fmt.Print("\nO arquivo ", numArq, " contém ", quantidade, " nota(s) ", nota)
						//fmt.Print("\nNas linhas:")
						//fmt.Print(linhas)

					} else {
						fmt.Print("\nO arquivo ", numArq, " não contém nenhuma nota ", nota)
					}

				} else {

					var linha string
					fmt.Print("\nQual a linha que deseja consultar? ")
					fmt.Scan(&linha)

					numLinha, _ := strconv.Atoi(linha)

					if tabela.BuscaBinariaNotaLinha(numArq, nota, numLinha) {
						fmt.Print("\nA nota ", nota, " existe na linha ", numLinha, " no arquivo ", numArq)
					} else {
						fmt.Print("\nA nota ", nota, " NÃO existe na linha ", numLinha, " no arquivo ", numArq)
					}

				}

				fmt.Print("\n\nDeseja continuar? ")
				fmt.Print("\n\n1 - Continuar")
				fmt.Print("\n2 - Sair")
				fmt.Print("\n\nInsira a opção que deseja: ")
				fmt.Scan(&opcao)

				if opcao == "2" {
					break
				}

			} else {
				fmt.Print("\n\nERRO: Nota inserida é inválida\n\n")
				break
			}

		} else {
			fmt.Print("\n\nERRO: Número do arquivo inserido é inválido\n\n")
			break
		}
	}

}
