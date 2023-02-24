package config

import (
	"github.com/gin-gonic/gin"
	"github.com/laioncorcino/go-user/controller"
)

func InitRoutes(router *gin.RouterGroup) {
	router.GET("/user/:userId", controller.FindById)
	router.GET("/user/:userEmail", controller.FindByEmail)
	router.POST("/user", controller.Create)
	router.PUT("/user/:userId", controller.Update)
	router.DELETE("/user/:userId", controller.Delete)
}
