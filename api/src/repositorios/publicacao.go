package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statment, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statment.Close()

	resultado, erro := statment.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)

	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Publicacoes) BuscarPublicacoes(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`SELECT DISTINCT p.*, u.nick FROM publicacoes AS p
			INNER join usuarios AS u ON u.id = p.autor_id
			INNER JOIN seguidores AS s ON p.autor_id = s.usuario_id
		WHERE u.id = ? OR s.seguidor_id = ? ORDER BY 1 DESC`,
		usuarioId, usuarioId,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CraidaEm,
			&publicacao.AutoNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) BuscarPublicacao(publicacaoId uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(
		`SELECT p.*, u.nick FROM publicacoes AS p
			INNER JOIN usuarios AS u ON (u.id = p.autor_id)
		WHERE p.id = ?`,
		publicacaoId,
	)

	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CraidaEm,
			&publicacao.AutoNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (repositorio Publicacoes) BuscarPublicacoesPorUsuario(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT p.*, u.nick FROM publicacoes AS p
			INNER JOIN usuarios AS u ON (u.id = p.autor_id)
		WHERE p.autor_id = ?
	`, usuarioId)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CraidaEm,
			&publicacao.AutoNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) AtualizarPublicacao(publicacaoId uint64, publicacao modelos.Publicacao) error {
	statment, erro := repositorio.db.Prepare(
		"UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) DeletarPublicacao(publicacaoId uint64) error {
	statment, erro := repositorio.db.Prepare(
		"DELETE FROM publicacoes WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) CurtirPublicacao(publicacaoId uint64) error {
	statment, erro := repositorio.db.Prepare(
		"UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro := statment.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) DescurtirPublicacao(publicacaoId uint64) error {
	statment, erro := repositorio.db.Prepare(
		`UPDATE publicacoes SET curtidas =
			CASE
				WHEN curtidas > 0 THEN curtidas - 1
				ELSE 0
			END
		WHERE id = ?`,
	)

	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro := statment.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}
