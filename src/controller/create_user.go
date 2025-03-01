package controller

import (
	"fmt"

	"github.com/Brun0Santos/api-go-testes/src/configuration/rest_err"
	"github.com/Brun0Santos/api-go-testes/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect field, errors=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
