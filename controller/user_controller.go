package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/laioncorcino/go-user/config/logger"
	"github.com/laioncorcino/go-user/json"
	"github.com/laioncorcino/go-user/service"
	"go.uber.org/zap"
)

func FindById(c *gin.Context) {

}

func FindByEmail(c *gin.Context) {

}

func Create(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var req json.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Error trying to marshal object", err, zap.String("journey", "createUser"))
		restErr := service.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
