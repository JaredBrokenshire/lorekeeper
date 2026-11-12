package routes

import (
	"lorekeeper/api"
	"lorekeeper/api/handlers"
	mw "lorekeeper/pkg/middleware"
	"os"
)

func fileRoutes(server *api.Server) {
	fileHandler := handlers.NewFileHandler(server)

	file := server.Echo.Group("/files")
	file.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	file.GET("/:filepath", fileHandler.Get)
}
