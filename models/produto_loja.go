package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Produto_cad_loja struct {
	Prlid                           uuid.UUID `gorm:"primaryKey" json:"id"`
	Prlid_empresa_cad               uuid.UUID `json:"id_empresa_cad"`
	Prlid_produto_cad               uuid.UUID `json:"id_produto_cad"`
	Conferencia                     bool      `json:"conferencia"`
	Prlestoque_minimo               float64   `json:"estoque_minimo"`
	Prlestoque_minimo_aviso         float64   `json:"estoque_minimo_aviso"`
	Prlcontrola_estoque             bool      `json:"controla_estoque"`
	Prlestoque_minimo_confere       float64   `json:"estoque_minimo_confere"`
	Prlestoque_minimo_aviso_confere float64   `json:"estoque_minimo_aviso_confere"`
	Prlmeta_conferencia             float64   `json:"meta_conferencia"`
	Prltipo_comissao                int       `json:"tipo_comissao"`
	Prlcomissao                     float64   `json:"comissao"`
	Prltipo_comissao_conferencia    int       `json:"tipo_comissao_conferencia"`
	Prlcomissao_conferencia         float64   `json:"comissao_conferencia"`
	Prlglp                          float64   `json:"glp"`
	Prlgnn                          float64   `json:"gnn"`
	Prlgni                          float64   `json:"gni"`
	Prlativo                        bool      `json:"ativo"`
	Prlid_ibge_estados_origem       uuid.UUID `json:"id_ibge_estados_origem"`
	Prlaltdata                      time.Time `json:"altdata"`
	Prlid_usuario                   uuid.UUID `json:"id_usuario"`
	Prlobriga_validade              bool      `json:"obriga_validade"`
	Prlbloqueia_custo_ent           bool      `json:"bloqueia_custo_ent"`

	//Prlpermite_compra               bool      `json:"permite_compra"`
	//Prlmargem_minima                float64   `json:"margem_minima"`
}

func ValidarDadosProdutoLoja(produto *Produto_cad_loja) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}

	return nil
}
