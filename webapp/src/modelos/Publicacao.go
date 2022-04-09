package modelos

import "time"

type Publicacao struct {
	Id        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas,omitempty"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}
