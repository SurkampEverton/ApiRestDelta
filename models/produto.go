package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Produto_cad struct {
	Prcid               uuid.UUID `gorm:"primaryKey" json:"id"`
	Prcnumero           int64     `json:"numero"`
	Prcproduto_tipo     int64     `json:"produto_tipo"`
	Prcid_produto_grupo uuid.UUID `json:"id_produto_grupo"`
	Prcnome             string    `json:"nome"`
	Prcnomereduzido     string    `json:"nomereduzido"`
	Prciat              string    `json:"iat"`
	Prcippt             string    `json:"ippt"`
	Prcespecificacao    string    `json:"especificacao"`
	Prcobs              string    `json:"obs"`
	Prcinfnutricional   string    `json:"infnutricional"`
	Prcexportar         string    `json:"exportar"`
	Prcorigem           string    `json:"origem"`
	Prcanp              string    `json:"anp"`
	Prcncm              string    `json:"ncm"`
	Prcativa_grade      bool      `json:"ativa_grade"`
	Prcpermite_venda    bool      `json:"permite_venda"`
	Prclista_venda      bool      `json:"lista_venda"`
	Prcativo            int       `json:"ativo"`
	Prcincdata          time.Time `json:"incdata"`
	Prcaltdata          time.Time `json:"altdata"`
	Prcid_usuario       uuid.UUID `json:"id_usuario"`
	Prcvalido           bool      `json:"valido"`
	Prcid_tributacao    string    `json:"id_tributacao"`
	Prcobriga_vendedor  int64     `json:"obriga_vendedor"`
	Prcid_imposto_ibpt  uuid.UUID `json:"id_imposto_ibpt"`
	Prcdecimais         int64     `json:"decimais"`

	//PrccodigoMensagem   int64     `json:"codigoMensagem"`
	//Prcestoqueminimo    int64     `json:"estoqueminimo"`
	//Prcid_externo       int64     `json:"id_externo"`
	//Prcfabricante       int64     `json:"fabricante"`
	//Prcid_lancamentospadroes_cad uuid.UUID `json:"id_lancamentospadroes_cad"`
	//prcfoto bytea,
	//Prchash            string    `json:"hash"`
	//Prcmargemlucro               float64   `json:"margemlucro"`
	//Prccomissao                  float64   `json:"comissao"`
}

func ValidarDadosProduto(produto *Produto_cad) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}

	return nil
}
