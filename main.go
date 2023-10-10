package main

import (
	"loja-web-go/routes"
	"net/http"
)

func main() {
	routes.Rotas()
	//inicia o servidor web e escuta a porta 8000
	http.ListenAndServe(":8000", nil)
}
