package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gothout/crud-go-study/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewNotFoundError("Voce chamou a rota de forma errada")
	c.JSON(err.Code, err)
}