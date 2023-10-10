package model

import "loja-web-go/db"

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {
	produtos := []Produto{}
	p := Produto{}
	bd := db.ConexaoBD()
	//realiza query select no BD e guarda em produtosBD
	produtosBD, err := bd.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	//percorre cada linha BD, converte na struct Produto e adiciona ao slice produtos
	for produtosBD.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = produtosBD.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}
	defer bd.Close()
	return produtos
}
