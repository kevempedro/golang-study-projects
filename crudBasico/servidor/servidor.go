package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário!"))
	}

	fmt.Println(usuario)

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	defer db.Close()

	statment, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")

	if erro != nil {
		w.Write([]byte("Erro ao preparar query de inserção dos dados!"))
		return
	}

	defer statment.Close()

	insercao, erro := statment.Exec(usuario.Nome, usuario.Email)

	if erro != nil {
		w.Write([]byte("Erro ao inserir os dados!"))
	}

	idInserido, erro := insercao.LastInsertId()

	if erro != nil {
		w.Write([]byte("Erro ao devolver o id do registro inserido!"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")

	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}

	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para json!"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmentro para inteiro!"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	defer db.Close()

	linha, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", ID)

	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}

	var usuario usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para json!"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	paramentros := mux.Vars(r)

	ID, erro := strconv.ParseUint(paramentros["id"], 10, 64)

	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
	}

	var usuario usuario

	if erro := json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conecatar no banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao preparar o statement!"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar banco de dados!"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
