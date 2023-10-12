package controllers

import (
	"log"
	"loja-web-go/model"
	"net/http"
	"strconv"
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

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		precoStr := r.FormValue("preco")
		quantidadeStr := r.FormValue("quantidade")

		preco, errP := strconv.ParseFloat(precoStr, 64)
		if errP != nil {
			log.Println("Erro na Conversão do Preço")
		}
		quantidade, errQ := strconv.Atoi(quantidadeStr)
		if errQ != nil {
			log.Println("Erro na Conversão da Quantidade")
		}

		model.AddProduto(nome, descricao, preco, quantidade)

		http.Redirect(w, r, "/", 301)
	}

}
