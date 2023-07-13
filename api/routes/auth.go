package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodolfoksveiga/k8s-go/handlers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)
	router.GET("/user", handlers.GetCurrentUser)
}
