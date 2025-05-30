package schemas

import "time"

type CustomerSchema struct {
	IdCliente     string    `db:"idCliente"`
	FlEmail       int       `db:"flEmail"`
	FlTwitch      int       `db:"flTwitch"`
	FlYouTube     int       `db:"flYouTube"`
	FlBlueSky     int       `db:"flBlueSky"`
	FlInstagram   int       `db:"flInstagram"`
	QtdePontos    int       `db:"qtdePontos"`
	DtCriacao     time.Time `db:"dtCriacao"`
	DtAtualizacao time.Time `db:"dtAtualizacao"`
}

type ProductSchema struct {
	IdProduto            string `db:"idProduto"`
	DescProduto          string `db:"descProduto"`
	DescCateogriaProduto string `db:"descCateogriaProduto"`
}

type TransactionProductSchema struct {
	IdTransacaoProduto string `db:"idTransacaoProduto"`
	IdTransacao        string `db:"idTransacao"`
	IdProduto          string `db:"idProduto"`
	QtdeProduto        int    `db:"qtdeProduto"`
	VlProduto          int    `db:"vlProduto"`
}

type TransactionSchema struct {
	IdTransacao       string    `db:"idTransacao"`
	IdCliente         string    `db:"idCliente"`
	DtCriacao         time.Time `db:"dtCriacao"`
	QtdePontos        int       `db:"qtdePontos"`
	DescSistemaOrigem string    `db:"descSistemaOrigem"`
}
