package controllers

import (
	"apirestdelta/database"
	"apirestdelta/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func PesquisaProdutoCompletoId(c *gin.Context) {
	var produto models.ProdutoCompleto

	id_loja := c.Param("id_loja")
	id := c.Param("id")

	database.DB.Where(&models.Produto_cad{
		Prcid: uuid.MustParse(id)}).Find(&produto.Produto)

	if produto.Produto.Prcid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto não encontrado para pesquisa completa",
		})

		return
	}

	database.DB.Where(&models.Produto_cad_loja{
		Prlid_empresa_cad: uuid.MustParse(id_loja),
		Prlid_produto_cad: uuid.MustParse(id),
	}).Find(&produto.ProdutoLoja)

	if produto.ProdutoLoja.Prlid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto loja não encontado para pesquisa completa",
		})

		return
	}

	database.DB.Model(&models.Produto_deposito{}).Where(
		"Prdid_empresa_cad = ? and Prdid_produto_cad = ?", uuid.MustParse(id_loja), uuid.MustParse(id),
	).Find(&produto.ProdutoDeposito)

	var ids_codigo []uuid.UUID

	database.DB.Where(&models.Produto_codigo{
		Prcdid_produto_cad: uuid.MustParse(id),
	}).Order("Prcdcodigo").Find(&produto.ProdutoCodigo)

	database.DB.Model(&models.Produto_codigo{}).Where(
		"Prcdid_produto_cad = ? ", uuid.MustParse(id),
	).Pluck("Prcdid", &ids_codigo)

	database.DB.Model(&models.Preco{}).Where(
		"Prcid_empresa_cad = ? and Prcid_produto_codigo in (?) ", uuid.MustParse(id_loja), ids_codigo,
	).Order("Prcatvdata desc").Find(&produto.Preco)

	database.DB.Where(&models.Produto_cad_composto{
		Prcid_produto_cad: uuid.MustParse(id),
		Prcstatus:         0,
	}).Find(&produto.ProdutoComposto)

	database.DB.Model(&models.Preco{}).Where(
		"Prcid_empresa_cad = ? and Prcid_produto_codigo in (?) ", uuid.MustParse(id_loja), ids_codigo,
	).Order("Prcatvdata desc").Find(&produto.Preco)

	c.JSON(http.StatusOK, produto)
}

func SalvaProdutoCompleto(c *gin.Context) {
	tx := database.DB.Begin()

	var produto_completo models.ProdutoCompleto

	if err := c.ShouldBindJSON(&produto_completo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := models.ValidarDadosProduto(&produto_completo.Produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := models.ValidarDadosProdutoLoja(&produto_completo.ProdutoLoja); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		tx.Rollback()
		return
	}

	/*if err := models.ValidarDadosProdutoDeposito(&produto_completo.ProdutoDeposito); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := models.ValidarDadosProdutoCodigo(&produto_completo.ProdutoCodigo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := models.ValidarDadosProdutoComposto(&produto_completo.ProdutoComposto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := models.ValidarDadosPreco(&produto_completo.Preco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		tx.Rollback()
		return
	}*/

	tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&produto_completo.Produto)
	tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&produto_completo.ProdutoLoja)

	if len(produto_completo.ProdutoDeposito) != 0 {
		tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&produto_completo.ProdutoDeposito)
	}

	if len(produto_completo.ProdutoCodigo) != 0 {
		tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&produto_completo.ProdutoCodigo)
	}

	if len(produto_completo.ProdutoComposto) != 0 {
		tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&produto_completo.ProdutoComposto)
	}

	if len(produto_completo.Preco) != 0 {
		tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&produto_completo.Preco)
	}

	tx.Commit()
	c.JSON(http.StatusOK, produto_completo)
}
