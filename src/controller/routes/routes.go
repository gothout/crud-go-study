package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gothout/crud-go-study/src/controller"
)

//Inicializando rotas
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.CreateUser)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleUser/:userId", controller.DeleteUser)
}