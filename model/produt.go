package model

import "loja-web-go/db"

type Produto struct {
	Id         int
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
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}
	defer bd.Close()
	return produtos
}

func AddProduto(nome, descricao string, preco float64, quantidade int) {
	bd := db.ConexaoBD()
	//prepara o banco para inserção do produto
	inserirBD, err := bd.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	//se não houver erro nas linhas anteriores, insere o produto no banco
	inserirBD.Exec(nome, descricao, preco, quantidade)
	defer bd.Close()
}

func DeletaProduto(id string) {
	bd := db.ConexaoBD()

	deletarBD, err := bd.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarBD.Exec(id)
	defer bd.Close()
}
