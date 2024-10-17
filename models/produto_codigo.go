package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Produto_codigo struct {
	Prcdid                 uuid.UUID `json:"Prcdid"`
	Prcdid_produto_cad     uuid.UUID `json:"Prcdid_produto_cad"`
	Prcdid_produto_cad_pai uuid.UUID `json:"Prcdid_produto_cad_pai"`
	Prcdid_unidades        uuid.UUID `json:"Prcdid_unidades"`
	Prcdcodigo             string    `json:"Prcdcodigo"`
	Prcdcodigobarra        string    `json:"Prcdcodigobarra"`
	Prcdcodigoref          string    `json:"Prcdcodigoref"`
	Prcdpeso_bruto         float64   `json:"Prcdpeso_bruto"`
	Prcdpeso_liquido       float64   `json:"Prcdpeso_liquido"`
	Prcdtaxa_conversao     float64   `json:"Prcdtaxa_conversao"`
	Prcdfiscal             bool      `json:"Prcdfiscal"`
	Prcdhash               string    `json:"Prcdhash"`
	Prcdincdata            time.Time `json:"Prcdincdata"`
	Prcdid_usuario         uuid.UUID `json:"Prcdid_usuario"`
	Prcdid_externo         int       `json:"Prcdid_externo"`
}

func ValidarDadosProdutoCodigo(produto *Produto_codigo) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}

	return nil
}
