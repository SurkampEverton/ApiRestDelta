package controllers

import (
	"apirestdelta/database"
	"apirestdelta/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TodosProdutosLoja(c *gin.Context) {
	var p []models.Produto_cad_loja
	database.DB.Find(&p)
	c.JSON(200, p)
}

func PesquisaProdutoLojaId(c *gin.Context) {
	var produto models.Produto_cad_loja
	id := c.Param("id")
	database.DB.Where(&models.Produto_cad_loja{
		Prlid: uuid.MustParse(id)}).First(&produto)

	if produto.Prlid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto loja não encontrada",
		})

		return
	}

	c.JSON(http.StatusOK, produto)
}

func PesquisaProdutoLojaProdutoId(c *gin.Context) {
	var produto models.Produto_cad_loja
	id := c.Param("id")
	database.DB.Where(&models.Produto_cad_loja{
		Prlid_produto_cad: uuid.MustParse(id)}).First(&produto)

	if produto.Prlid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto loja não encontrada",
		})

		return
	}

	c.JSON(http.StatusOK, produto)
}

func NovoProdutoLoja(c *gin.Context) {
	var produto models.Produto_cad_loja

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosProdutoLoja(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&produto)
	c.JSON(http.StatusOK, produto)
}

func EditaProdutoLoja(c *gin.Context) {
	var produto models.Produto_cad_loja

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosProdutoLoja(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&produto).UpdateColumns(produto)
	c.JSON(http.StatusOK, produto)
}

func DeletarProdutoLoja(c *gin.Context) {
	var produto models.Produto_cad_loja
	id := c.Params.ByName("id")
	database.DB.Delete(&produto, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Produto loja deletado com sucesso",
	})
}
