package leitura

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Dado struct {
	Arquivo int    `json:"arq"`
	Ordem   int    `json:"ordem"`
	Notas   string `json:"notas"`
}

func verifErro(erro error) {

	if erro != nil {
		fmt.Println(erro)
	}

}

func LeituraJson(arquivo string) []Dado {

	dados := make([]Dado, 0)

	arqJson, erro := os.Open(arquivo)
	verifErro(erro)

	defer arqJson.Close()

	bytesJson, erro := ioutil.ReadAll(arqJson)
	verifErro(erro)

	erro = json.Unmarshal(bytesJson, &dados)
	verifErro(erro)

	return dados

}
