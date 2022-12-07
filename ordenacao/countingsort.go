package ordenacao

import (
	arq "BuscaMusica/leituraescrita"
)

func encontraMaior(dados []arq.Dado) int {

	maior := dados[0].Ordem

	for _, d := range dados {

		if maior < d.Ordem {
			maior = d.Ordem
		}

	}

	return maior

}

func iniciaVetor(maior int) []int {

	vetorContagem := make([]int, maior+1)

	for i := range vetorContagem {
		vetorContagem[i] = 0
	}

	return vetorContagem

}

func CountingSort(dados []arq.Dado) []arq.Dado {

	vetorContagem := iniciaVetor(encontraMaior(dados))

	for _, d := range dados {
		vetorContagem[d.Ordem]++
	}

	for i := 1; i < len(vetorContagem); i++ {
		vetorContagem[i] += vetorContagem[i-1]
	}

	vetorResultado := make([]arq.Dado, len(dados))

	for i := 0; i < len(dados); i++ {

		d := dados[i]
		pos := vetorContagem[d.Ordem] - 1

		vetorResultado[pos] = d
		vetorContagem[d.Ordem] = vetorContagem[d.Ordem] - 1

	}

	return vetorResultado

}
