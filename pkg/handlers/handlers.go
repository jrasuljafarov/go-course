package handlers

import (
	"net/http"
)

func About(writer http.ResponseWriter, request *http.Request) {
	RenderTemplate(writer, "about.page.tmpl")
}

func Home(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "home.page.tmpl")
}
