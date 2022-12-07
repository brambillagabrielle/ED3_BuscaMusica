package tabelahash

import (
	busca "BuscaMusica/busca"
	arq "BuscaMusica/leituraescrita"
	mapa "BuscaMusica/mapa"
	"unicode/utf8"
)

type TabelaHash struct {
	notas map[int][]arq.Dado
}

func hash(numArq int) (indice int) {
	return mapa.RetornaIndiceMapa(numArq)
}

func (tabHash *TabelaHash) FormaTabelaHash(dados []arq.Dado) {

	for _, d := range dados {
		tabHash.Inserir(d)
	}

}

func (tabHash *TabelaHash) Inserir(dado arq.Dado) {

	if tabHash.notas == nil {
		tabHash.notas = make(map[int][]arq.Dado)
	}

	indice := hash(dado.Arquivo)
	tabHash.notas[indice] = append(tabHash.notas[indice], dado)

}

func (tabHash *TabelaHash) BuscarNotaArquivo(numArq int, nota string) (quantidade int, linhas []int) {

	indice := hash(numArq)
	arquivo := tabHash.notas[indice]
	runeNota, _ := utf8.DecodeRuneInString(nota)

	index := 0
	quantidade = 0
	linhas = []int{}
	for _, a := range arquivo {

		for _, c := range a.Notas {

			if c == runeNota {

				quantidade++

				if quantidade == 1 {

					index++
					linhas = append(linhas, a.Ordem)

				} else if quantidade > 1 {

					if linhas[index-1] != a.Ordem {

						index++
						linhas = append(linhas, a.Ordem)

					}

				}

			}

		}

	}

	return quantidade, linhas

}

func (tabHash *TabelaHash) BuscarNotaLinha(numArq int, nota string, numLinha int) (existe bool) {

	indice := hash(numArq)
	arquivo := tabHash.notas[indice]
	runeNota, _ := utf8.DecodeRuneInString(nota)

	for _, a := range arquivo {

		if a.Ordem == numLinha {

			for _, c := range a.Notas {

				if c == runeNota {
					return true
				}

			}

		}

	}

	return false

}

func (tabHash *TabelaHash) BuscaBinariaNotaLinha(numArq int, nota string, numLinha int) (existe bool) {

	indice := hash(numArq)
	arquivo := tabHash.notas[indice]
	runeNota, _ := utf8.DecodeRuneInString(nota)

	index := busca.BuscaBinaria(arquivo, numLinha)

	if indice == -1 {
		return false
	} else {

		for _, c := range arquivo[index].Notas {

			if c == runeNota {
				return true
			}

		}

	}

	return false

}
