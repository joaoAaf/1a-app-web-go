package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// conexão com o banco de dados
func ConexaoBD() *sql.DB {
	//string de conexão
	conexao := "user=postgres dbname=loja_golang password=123456 host=localhost sslmode=disable"
	//abre uma conexão com o banco
	bd, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return bd
}
