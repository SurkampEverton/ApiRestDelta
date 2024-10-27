package models

import (
	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Produto_deposito struct {
	Prdid             uuid.UUID `gorm:"primaryKey" json:"id"`
	Prdid_empresa_cad uuid.UUID `json:"id_empresa_cad"`
	Prdid_deposito    uuid.UUID `json:"id_deposito"`
	Prdid_produto_cad uuid.UUID `json:"id_produto_cad"`
	Prdprincipal      bool      `json:"principal"`
	Prdstatus         int       `json:"status"`
}

func ValidarDadosProdutoDeposito(produto *Produto_deposito) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}

	return nil
}
