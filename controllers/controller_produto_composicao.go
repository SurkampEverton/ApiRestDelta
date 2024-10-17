package controllers

import (
	"apirestdelta/database"
	"apirestdelta/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TodosProdutosComposto(c *gin.Context) {
	var p []models.Produto_cad_composto
	database.DB.Find(&p)
	c.JSON(200, p)
}

func PesquisaProdutoCompostoId(c *gin.Context) {
	var produto models.Produto_cad_composto
	id := c.Param("id")
	database.DB.Where(&models.Produto_cad_composto{
		Prccid: uuid.MustParse(id)}).First(&produto)

	if produto.Prccid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto loja não encontrada",
		})

		return
	}

	c.JSON(http.StatusOK, produto)
}

func PesquisaProdutoCompostoProdutoId(c *gin.Context) {
	var produto models.Produto_cad_composto
	id := c.Param("id")
	database.DB.Where(&models.Produto_cad_composto{
		Prcid_produto_cad: uuid.MustParse(id)}).First(&produto)

	if produto.Prccid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto loja não encontrada",
		})

		return
	}

	c.JSON(http.StatusOK, produto)
}

func NovoProdutoComposto(c *gin.Context) {
	var produto models.Produto_cad_composto

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosProdutoComposto(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&produto)
	c.JSON(http.StatusOK, produto)
}

func EditaProdutoComposto(c *gin.Context) {
	var produto models.Produto_cad_composto

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosProdutoComposto(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&produto).UpdateColumns(produto)
	c.JSON(http.StatusOK, produto)
}

func DeletarProdutoComposto(c *gin.Context) {
	var produto models.Produto_cad_composto
	id := c.Params.ByName("id")
	database.DB.Delete(&produto, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Produto loja deletado com sucesso",
	})
}
