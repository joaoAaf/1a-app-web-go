package routes

import (
	"loja-web-go/controllers"
	"net/http"
)

func Rotas() {
	//vicula a função index a rota "/"
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/register", controllers.Register)
}
