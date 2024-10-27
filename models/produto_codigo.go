package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Produto_codigo struct {
	Prcdid             uuid.UUID `gorm:"primaryKey" json:"id"`
	Prcdid_produto_cad uuid.UUID `json:"id_produto_cad"`
	Prcdid_unidades    uuid.UUID `json:"id_unidades"`
	Prcdcodigo         string    `json:"codigo"`
	Prcdcodigobarra    string    `json:"codigobarra"`
	Prcdcodigoref      string    `json:"codigoref"`
	Prcdpeso_bruto     float64   `json:"peso_bruto"`
	Prcdpeso_liquido   float64   `json:"peso_liquido"`
	Prcdtaxa_conversao float64   `json:"taxa_conversao"`
	Prcdfiscal         bool      `json:"fiscal"`
	Prcdincdata        time.Time `json:"incdata"`
	Prcdid_usuario     uuid.UUID `json:"id_usuario"`
}

func ValidarDadosProdutoCodigo(produto *Produto_codigo) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}

	return nil
}
