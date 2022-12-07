package busca

import (
	arq "BuscaMusica/leituraescrita"
)

func BuscaBinaria(dados []arq.Dado, linha int) int {

	menor := 0
	maior := len(dados) - 1
	var media int

	for menor <= maior {

		media = (menor + maior) / 2

		if dados[media].Ordem < linha {
			menor = media + 1
		} else {
			maior = media - 1
		}

	}

	if menor == len(dados) || dados[menor].Ordem != linha {
		return -1
	}

	return media

}
