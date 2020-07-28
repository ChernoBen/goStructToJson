package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	//Criando mapeamento p/ exportar para json
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct to json
	p1 := produto{1, "Laptop", 1000, []string{"AMD", "Ryzen"}}
	//json.Marshal retorna o json e um erro
	p1Json, _ := json.Marshal(p1)
	// Como p1Json possui um slice de bytes devemos converter p/ string para poder imprimir no terminal
	fmt.Println(string(p1Json))

	//agora fazendo o metodo inverso, json p/ struct
	var p2 produto // estura declarada mas nao instanciada
	jsonString := `{"id":2,"nome":"Acer Aspire","preco":1665.80,"tags":["AMD","Ryzen 5"]}`
	// json carregado na varivel agora criar instanciar em uma estrutura
	json.Unmarshal([]byte(jsonString), &p2) // esse metodo pede um slice de bites contendo o json e o endereco de memoria da esturura que queremos instanciar os dados do json
	// testando
	fmt.Println(p2)

}
