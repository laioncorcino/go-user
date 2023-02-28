package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/laioncorcino/go-user/config/logger"
	"github.com/laioncorcino/go-user/json"
	"github.com/laioncorcino/go-user/service"
	"github.com/laioncorcino/go-user/validator"
	"go.uber.org/zap"
	"net/http"
)

var (
	userService service.IUserService
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
		restErr := validator.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	user, err := userService.CreateUser(req)
	if err != nil {
		logger.Error("Error trying to call CreateUser service", err, zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", user.ID), zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
