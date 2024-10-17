package controllers

import (
	"apirestdelta/database"
	"apirestdelta/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TodosPreco(c *gin.Context) {
	var p []models.Preco
	database.DB.Find(&p)
	c.JSON(200, p)
}

func PesquisaPrecoId(c *gin.Context) {
	var produto models.Preco
	id := c.Param("id")
	database.DB.Where(&models.Preco{
		Prcid: uuid.MustParse(id)}).First(&produto)

	if produto.Prcid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto loja não encontrada",
		})

		return
	}

	c.JSON(http.StatusOK, produto)
}

func PesquisaPrecoProdutoCodigoId(c *gin.Context) {
	var produto models.Preco
	id := c.Param("id")
	database.DB.Where(&models.Preco{
		Prcid_produto_codigo: uuid.MustParse(id)}).First(&produto)

	if produto.Prcid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto loja não encontrada",
		})

		return
	}

	c.JSON(http.StatusOK, produto)
}

func NovoPreco(c *gin.Context) {
	var produto models.Preco

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosPreco(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&produto)
	c.JSON(http.StatusOK, produto)
}

func EditaPreco(c *gin.Context) {
	var produto models.Preco

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosPreco(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&produto).UpdateColumns(produto)
	c.JSON(http.StatusOK, produto)
}

func DeletarPreco(c *gin.Context) {
	var produto models.Preco
	id := c.Params.ByName("id")
	database.DB.Delete(&produto, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Produto loja deletado com sucesso",
	})
}
