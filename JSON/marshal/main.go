package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	c := cachorro{"Rex", "Dálmata", 6}

	fmt.Println(c)

	cahorroJSON, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cahorroJSON)
	fmt.Println(bytes.NewBuffer(cahorroJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cahorro2JSON, erro := json.Marshal(c2)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cahorro2JSON))
}

type cachorro struct {
	None  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}
