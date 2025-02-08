package controller

import (
	"github.com/Brun0Santos/api-go-testes/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func FindUserByID(c *gin.Context) {
	err := rest_err.NewBadRequestError("error no findbyid")
	c.JSON(err.Code, err)
}

func FindUserByEmail(c *gin.Context) {

}
