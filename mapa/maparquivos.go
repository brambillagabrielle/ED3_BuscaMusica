package mapa

import (
	arq "BuscaMusica/leituraescrita"
)

var MapaNumArq = make(map[int]int)

func addNumMapa(numeroArq int, indice int) {
	MapaNumArq[numeroArq] = indice
}

func RetornaIndiceMapa(numeroArq int) (indice int) {
	return MapaNumArq[numeroArq]
}

func CriaMapaArquivos(dados []arq.Dado) {

	ordem := dados[0].Ordem

	for i, d := range dados {

		if d.Ordem == ordem {
			addNumMapa(d.Arquivo, i)
		} else {
			break
		}

	}

}
