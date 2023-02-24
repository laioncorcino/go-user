package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/laioncorcino/go-user/json"
	"github.com/laioncorcino/go-user/rest_error"
)

func FindById(c *gin.Context) {

}

func FindByEmail(c *gin.Context) {

}

func Create(c *gin.Context) {
	var req json.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := rest_error.NewBadRequestError("Campos incorretos")
		c.JSON(restErr.Code, restErr)
		return
	}
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
