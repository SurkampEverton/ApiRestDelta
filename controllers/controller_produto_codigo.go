package controllers

import (
	"apirestdelta/database"
	"apirestdelta/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TodosProdutosCodigo(c *gin.Context) {
	var p []models.Produto_codigo
	database.DB.Find(&p)
	c.JSON(200, p)
}

func PesquisaProdutoCodigoId(c *gin.Context) {
	var produto models.Produto_codigo
	id := c.Param("id")
	database.DB.Where(&models.Produto_codigo{
		Prcdid: uuid.MustParse(id)}).First(&produto)

	if produto.Prcdid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Código do produto não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, produto)
}

func PesquisaProdutoCodigoProdutoId(c *gin.Context) {
	var produto models.Produto_codigo
	id := c.Param("id")
	database.DB.Where(&models.Produto_codigo{
		Prcdid_produto_cad: uuid.MustParse(id)}).First(&produto)

	if produto.Prcdid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Código do produto não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, produto)
}

func PesquisaProdutoCodigoPorCodigo(c *gin.Context) {
	var produto models.Produto_codigo
	codigo := c.Param("codigo")
	database.DB.Where(&models.Produto_codigo{
		Prcdcodigobarra: codigo}).First(&produto)

	if produto.Prcdid == uuid.Nil {
		database.DB.Where(&models.Produto_codigo{
			Prcdcodigobarra: codigo}).First(&produto)

		if produto.Prcdid == uuid.Nil {
			c.JSON(http.StatusNotFound, gin.H{
				"Not found": "Código do produto não encontrado",
			})

			return
		}
	}

	c.JSON(http.StatusOK, produto)
}

func NovoProdutoCodigo(c *gin.Context) {
	var produto models.Produto_codigo

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosProdutoCodigo(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&produto)
	c.JSON(http.StatusOK, produto)
}

func EditaProdutoCodigo(c *gin.Context) {
	var produto models.Produto_codigo

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosProdutoCodigo(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&produto).UpdateColumns(produto)
	c.JSON(http.StatusOK, produto)
}

func DeletarProdutoCodigo(c *gin.Context) {
	var produto models.Produto_codigo
	id := c.Params.ByName("id")
	database.DB.Delete(&produto, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Código do produto deletado com sucesso",
	})
}
