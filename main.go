package main

import (
	"github.com/joaoino100/api-go-gin/database"
	"github.com/joaoino100/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
