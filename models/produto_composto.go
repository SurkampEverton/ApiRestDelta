package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Produto_cad_composto struct {
	Prccid                   uuid.UUID `json:"Prcdid"`
	Prcid_produto_cad_origem uuid.UUID `json:"Prcid_produto_cad_origem"`
	Prcid_produto_cad        uuid.UUID `json:"Prcid_produto_cad"`
	Prcquantidade            float64   `json:"Prcquantidade"`
	Prcincdata               time.Time `json:"Prcincdata"`
	Prcaltdata               time.Time `json:"Prcaltdata"`
	Prcid_usuario            uuid.UUID `json:"Prcid_usuario"`
	Prcgrupo                 int       `json:"Prcgrupo"`
	Prcid_produto_codigo     uuid.UUID `json:"Prcid_produto_codigo"`
}

func ValidarDadosProdutoComposto(produto *Produto_cad_composto) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}

	return nil
}
