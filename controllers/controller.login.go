package controllers

import (
	"apirestdelta/database"
	"apirestdelta/models"
	"apirestdelta/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FazerLogin(c *gin.Context) {
	var login models.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidarDadosLogin(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var usuario models.Usuarios

	database.DB.Where(&models.Usuarios{
		Usnome: login.User}).First(&usuario)

	if usuario.Usid == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not found": "Usuário ou senha inválido",
		})

		return
	}

	valid := service.ComparePassword(login.Pass, usuario.Email)

	if valid != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not found": "Usuário ou senha inválido",
		})
	}

	token, err := service.GenerateToken(usuario.Usid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Not found": "Erro ao gerar token de para o usuário",
		})
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
