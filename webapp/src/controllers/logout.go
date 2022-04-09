package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)

	http.Redirect(w, r, "/login", 302)
}
