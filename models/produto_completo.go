package models

type ProdutoCompleto struct {
	Produto         Produto_cad            `json:"produto"`
	ProdutoLoja     Produto_cad_loja       `json:"produto_loja"`
	ProdutoDeposito []Produto_deposito     `json:"produto_deposito"`
	ProdutoCodigo   []Produto_codigo       `json:"produto_codigo"`
	ProdutoComposto []Produto_cad_composto `json:"produto_composto"`
	Preco           []Preco                `json:"preco"`
}
