package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Produto_cad_loja struct {
	Prlid                           uuid.UUID `json:"Prcdid"`
	Prlid_empresa_cad               uuid.UUID `json:"Prlid_empresa_cad"`
	Prlid_produto_cad               uuid.UUID `json:"Prlid_produto_cad"`
	Conferencia                     bool      `json:"Conferencia"`
	Prlestoque_minimo               float64   `json:"Prlestoque_minimo"`
	Prlestoque_minimo_aviso         float64   `json:"Prlestoque_minimo_aviso"`
	Prlcontrola_estoque             bool      `json:"Prlcontrola_estoque"`
	Prlestoque_minimo_confere       float64   `json:"Prlestoque_minimo_confere"`
	Prlestoque_minimo_aviso_confere float64   `json:"Prlestoque_minimo_aviso_confere"`
	Prlmeta_conferencia             float64   `json:"Prlmeta_conferencia"`
	Prltipo_comissao                int       `json:"Prltipo_comissao"`
	Prlcomissao                     float64   `json:"Prlcomissao"`
	Prltipo_comissao_conferencia    int       `json:"Prltipo_comissao_conferencia"`
	Prlcomissao_conferencia         float64   `json:"Prlcomissao_conferencia"`
	Prlglp                          float64   `json:"Prlglp"`
	Prlgnn                          float64   `json:"Prlgnn"`
	Prlgni                          float64   `json:"Prlgni"`
	Prlativo                        bool      `json:"Prlativo"`
	Prlid_ibge_estados_origem       uuid.UUID `json:"Prlid_ibge_estados_origem"`
	Prlaltdata                      time.Time `json:"Prlaltdata"`
	Prlid_usuario                   uuid.UUID `json:"Prlid_usuario"`
	Prlpermite_compra               bool      `json:"Prlpermite_compra"`
	Prlmargem_minima                float64   `json:"Prlmargem_minima"`
	Prlobriga_validade              bool      `json:"Prlobriga_validade"`
	Prlbloqueia_custo_ent           bool      `json:"Prlbloqueia_custo_ent"`
}

func ValidarDadosProdutoLoja(produto *Produto_cad_loja) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}

	return nil
}
