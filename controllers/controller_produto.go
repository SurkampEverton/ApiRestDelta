package controllers

import (
	"apirestdelta/database"
	"apirestdelta/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TodosProdutos(c *gin.Context) {
	var p []models.Produto_cad
	database.DB.Find(&p)
	c.JSON(200, p)
}

func PesquisaProdutoId(c *gin.Context) {
	var produto models.Produto_cad
	id := c.Param("id")
	database.DB.Where(&models.Produto_cad{
		Prcid: uuid.MustParse(id)}).First(&produto)

	if produto.Prcid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, produto)
}

func NovoProduto(c *gin.Context) {
	var produto models.Produto_cad

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosProduto(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&produto)
	c.JSON(http.StatusOK, produto)
}

func EditaProduto(c *gin.Context) {
	var produto models.Produto_cad

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosProduto(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&produto).UpdateColumns(produto)
	c.JSON(http.StatusOK, produto)
}

func DeletarProduto(c *gin.Context) {
	var produto models.Produto_cad
	id := c.Params.ByName("id")
	database.DB.Delete(&produto, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Produto deletado com sucesso",
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
