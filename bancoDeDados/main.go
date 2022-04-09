package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringUrlConexao := "root:root@/golang_study?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringUrlConexao)

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}
