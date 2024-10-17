package controllers

import (
	"apirestdelta/database"
	"apirestdelta/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PesquisaProdutoCompletoId(c *gin.Context) {
	var produto models.Produto_cad
	var produto_loja models.Produto_cad_loja

	id := c.Param("id")
	database.DB.Where(&models.Produto_cad{
		Prcid: uuid.MustParse(id)}).Find(&produto, &produto_loja)

	if produto.Prcid == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto completo não encontrada",
		})

		return
	}

	c.JSON(http.StatusOK, produto)
}
