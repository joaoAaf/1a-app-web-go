package controllers

import (
	"loja-web-go/model"
	"net/http"
	"text/template"
)

// cria uma varivel para receber as páginas html da pasta templates
var tpl = template.Must(template.ParseGlob("templates/*.html"))

// vincula a tabela de produtos a página index.html
func Index(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaProdutos()
	tpl.ExecuteTemplate(w, "index", produtos)
}

func Register(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register", nil)
}
