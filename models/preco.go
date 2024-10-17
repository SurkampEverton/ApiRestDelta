package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Preco struct {
	Prcid                   uuid.UUID `json:"Prcdid"`
	Prcid_empresa_cad       uuid.UUID `json:"Prcid_empresa_cad"`
	Prcid_produto_cad       uuid.UUID `json:"Prcid_produto_cad"`
	Prcid_preco_tabela      uuid.UUID `json:"Prcid_preco_tabela"`
	Prcvalor                float64   `json:"Prcvalor"`
	Prcvalido               bool      `json:"Prcvalido"`
	Prcincdata              time.Time `json:"Prcincdata"`
	Prcaltdata              time.Time `json:"Prcaltdata"`
	Prcid_usuario           uuid.UUID `json:"Prcid_usuario"`
	Prcid_produto_codigo    uuid.UUID `json:"Prcid_produto_codigo"`
	Prcatvdata              time.Time `json:"Prcatvdata"`
	Prcid_movimentoprodutos uuid.UUID `json:"Prcid_movimentoprodutos"`
	Prcagendamento          bool      `json:"Prcagendamento"`
}

func ValidarDadosPreco(preco *Preco) error {
	if err := validator.Validate(preco); err != nil {
		return err
	}

	return nil
}
