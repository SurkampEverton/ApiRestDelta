package main

import (
	"apirestdelta/database"
	"apirestdelta/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
