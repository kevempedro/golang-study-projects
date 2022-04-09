package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Kevem",
		"sobrenome": "Lima",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Kevem",
			"segundo":  "Lima",
		},

		"curso": {
			"nome": "Sistemas de Informaçãp",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Leão",
	}

	fmt.Println(usuario2)
}
