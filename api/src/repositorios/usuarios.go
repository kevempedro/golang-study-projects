package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint(ultimoIdInserido), nil
}

func (repositorio Usuarios) BuscarPorNomeOuNick(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome like ? OR nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorId(id uint64) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?",
		id,
	)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"SELECT id, senha FROM usuarios WHERE email = ?",
		email,
	)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Atualizar(id uint64, usuario modelos.Usuario) error {
	statment, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) Deletar(id uint64) error {
	statment, erro := repositorio.db.Prepare(
		"DELETE FROM usuarios WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) SeguirUsuario(usuarioId, seguidorId uint64) error {
	statment, erro := repositorio.db.Prepare(
		"INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)",
	)

	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) PararDeSeguirUsuario(usuarioId, seguidorId uint64) error {
	statment, erro := repositorio.db.Prepare(
		"DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarSeguidores(usuarioId uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email, u.criadoEm FROM usuarios AS u
			INNER JOIN seguidores AS s ON (u.id = s.seguidor_id)
		WHERE s.usuario_id = ?
	`, usuarioId)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarSeguindo(usuarioId uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT u.id, u.nome, u.nick, u.email, u.criadoEm FROM usuarios AS u
		INNER JOIN seguidores AS s ON (u.id = s.usuario_id)
	WHERE s.seguidor_id = ?
	`, usuarioId)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linha, erro := repositorio.db.Query("SELECT senha FROM usuarios WHERE id = ?", usuarioId)

	if erro != nil {
		return "", erro
	}

	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

func (repositorio Usuarios) AtualizarSenha(usuarioId uint64, novaSenhaComHash string) error {
	statment, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET senha = ? WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(novaSenhaComHash, usuarioId); erro != nil {
		return erro
	}

	return nil
}
