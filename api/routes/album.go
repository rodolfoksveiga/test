package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodolfoksveiga/k8s-go/handlers"
	"github.com/rodolfoksveiga/k8s-go/middlewares"
)

func AlbumRoutes(router *gin.Engine) {
	protectedRouter := router.Use(middlewares.JwtAuthMiddleware())

	protectedRouter.GET("/albums", handlers.GetAlbums)
	protectedRouter.GET("/albums/:id", handlers.GetAlbumById)
	protectedRouter.POST("/albums", handlers.CreateAlbum)
	protectedRouter.DELETE("/albums/:id", handlers.DeleteAlbum)
	protectedRouter.PUT("/albums/:id", handlers.UpdateAlbum)
}
