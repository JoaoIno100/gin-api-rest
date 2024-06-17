package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoino100/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.Run(":8080")
}
