package controllers

import (
	"apirestdelta/database"
	"apirestdelta/models"
	"apirestdelta/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditaUsuario(c *gin.Context) {
	var usuario models.Usuarios

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosUsuario(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	pass, err := service.EncryptPassword(usuario.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	usuario.Email = pass

	database.DB.Model(&usuario).UpdateColumns(usuario)
	c.Status(http.StatusOK)
}
