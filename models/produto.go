package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Produto_cad struct {
	Prcid                        uuid.UUID `json:"Prcid"`
	Prcnumero                    int64     `json:"Prcnumero"`
	Prcproduto_tipo              int64     `json:"Prcproduto_tipo"`
	Prcid_produto_grupo          uuid.UUID `json:"Prcid_produto_grupo"`
	Prcnome                      string    `json:"Prcnome"`
	Prcnomereduzido              string    `json:"Prcnomereduzido"`
	Prcid_lancamentospadroes_cad uuid.UUID `json:"Prcid_lancamentospadroes_cad"`
	Prciat                       string    `json:"Prciat"`
	Prcippt                      string    `json:"Prcippt"`
	Prcmargemlucro               float64   `json:"Prcmargemlucro"`
	Prccomissao                  float64   `json:"Prccomissao"`
	Prcespecificacao             string    `json:"Prcespecificacao"`
	Prcobs                       string    `json:"Prcobs"`
	Prcestoqueminimo             int64     `json:"Prcestoqueminimo"`
	Prcinfnutricional            string    `json:"Prcinfnutricional"`
	Prcfabricante                int64     `json:"Prcfabricante"`
	Prcexportar                  string    `json:"Prcexportar"`
	Prcorigem                    string    `json:"Prcorigem"`
	PrccodigoMensagem            int64     `json:"PrccodigoMensagem"`
	Prcanp                       string    `json:"Prcanp"`
	Prcncm                       string    `json:"Prcncm"`
	//prcfoto bytea,
	Prcativa_grade     bool      `json:"Prcativa_grade"`
	Prcpermite_venda   bool      `json:"Prcpermite_venda"`
	Prclista_venda     bool      `json:"Prclista_venda"`
	Prcativo           int       `json:"Prcativo"`
	Prchash            string    `json:"Prchash"`
	Prcincdata         time.Time `json:"Prcincdata"`
	Prcaltdata         time.Time `json:"Prcaltdata"`
	Prcid_usuario      uuid.UUID `json:"Prcid_usuario"`
	Prcvalido          bool      `json:"Prcvalido"`
	Prcid_tributacao   string    `json:"Prcid_tributacao"`
	Prcobriga_vendedor int64     `json:"Prcobriga_vendedor"`
	Prcid_imposto_ibpt uuid.UUID `json:"Prcid_imposto_ibpt"`
	Prcid_externo      int64     `json:"Prcid_externo"`
	Prcdecimais        int64     `json:"Prcdecimais"`
}

func ValidarDadosProduto(produto *Produto_cad) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}

	return nil
}
