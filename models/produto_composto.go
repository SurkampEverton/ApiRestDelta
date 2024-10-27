package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Produto_cad_composto struct {
	Prccid               uuid.UUID `gorm:"primaryKey" json:"id"`
	Prcid_produto_codigo uuid.UUID `json:"id_produto_codigo"`
	Prcid_produto_cad    uuid.UUID `json:"id_produto_cad"`
	Prcquantidade        float64   `json:"quantidade"`
	Prcincdata           time.Time `json:"incdata"`
	Prcaltdata           time.Time `json:"altdata"`
	Prcid_usuario        uuid.UUID `json:"id_usuario"`
	Prcgrupo             int       `json:"grupo"`
	Prcstatus            int       `json:"status"`
}

func ValidarDadosProdutoComposto(produto *Produto_cad_composto) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}

	return nil
}
