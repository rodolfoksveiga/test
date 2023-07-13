package main

import (
	"github.com/rodolfoksveiga/k8s-go/config"
	"github.com/rodolfoksveiga/k8s-go/routes"
)

func bbbbb() {
	config.ConnectDB()
	config.MigrateDB()
}

func main() {
	router := routes.SetupRouter()

	routes.AuthRoutes(router)
	routes.AlbumRoutes(router)

	router.Run()
}
