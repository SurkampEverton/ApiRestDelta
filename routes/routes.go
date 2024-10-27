package routes

import (
	"apirestdelta/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	/*
		r.GET("/produto", controllers.TodosProdutos)
		r.POST("produto", controllers.NovoProduto)
		r.GET("/produto/:id", controllers.PesquisaProdutoId)
		r.DELETE("/produto/:id", controllers.DeletarProduto)
		r.PATCH("/produto/:id", controllers.EditaProduto)
	*/

	r.GET("/produtocompleto/:id_loja/:id", controllers.PesquisaProdutoCompletoId)
	r.POST("produtocompleto", controllers.SalvaProdutoCompleto)

	r.NoRoute(controllers.RotaNaoEncontrada)

	r.Run()
}
