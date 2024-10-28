package routes

import (
	"apirestdelta/controllers"
	"apirestdelta/service"

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

	r.POST("/usuario/:id", controllers.EditaUsuario)
	r.POST("/login", controllers.FazerLogin)
	r.NoRoute(controllers.RotaNaoEncontrada)

	public := r.Group("/api")

	public.POST("/produto_completo/:id", controllers.EditaUsuario)
	public.POST("/login", controllers.FazerLogin)

	protected := r.Group("/api/admin")
	protected.Use(service.JwtAuthMiddleware())
	protected.POST("produto_completo", controllers.SalvaProdutoCompleto)
	protected.GET("/produto_completo/:id_loja/:id", controllers.PesquisaProdutoCompletoId)

	r.Run(":8080")

	r.Run()
}
