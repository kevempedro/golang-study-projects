package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(metodo, url, dados)

	if erro != nil {
		return nil, erro
	}

	cookie, _ := cookies.Ler(r)

	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	cliente := &http.Client{}

	response, erro := cliente.Do(request)

	if erro != nil {
		return nil, erro
	}

	return response, nil
}
