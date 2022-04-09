package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("App Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{
		"Kevem",
		"kevem@gmail.com",
	}

	templates.ExecuteTemplate(w, "home.html", u)
}

type usuario struct {
	Nome  string
	email string
}
