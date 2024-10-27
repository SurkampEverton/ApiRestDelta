package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Preco struct {
	Prcid                   uuid.UUID `gorm:"primaryKey" json:"id"`
	Prcid_empresa_cad       uuid.UUID `json:"id_empresa_cad"`
	Prcid_produto_cad       uuid.UUID `json:"id_produto_cad"`
	Prcid_preco_tabela      uuid.UUID `json:"id_preco_tabela"`
	Prcvalor                float64   `json:"valor"`
	Prcvalido               bool      `json:"valido"`
	Prcincdata              time.Time `json:"incdata"`
	Prcaltdata              time.Time `json:"altdata"`
	Prcid_usuario           uuid.UUID `json:"id_usuario"`
	Prcid_produto_codigo    uuid.UUID `json:"id_produto_codigo"`
	Prcatvdata              time.Time `json:"atvdata"`
	Prcid_movimentoprodutos uuid.UUID `json:"id_movimentoprodutos"`
	Prcagendamento          bool      `json:"agendamento"`
}

func ValidarDadosPreco(preco *Preco) error {
	if err := validator.Validate(preco); err != nil {
		return err
	}

	return nil
}
