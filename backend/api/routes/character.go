package routes

import (
	"lorekeeper/api"
	"lorekeeper/api/handlers"
	mw "lorekeeper/pkg/middleware"
	"os"
)

func characterRoutes(s *api.Server) {
	characterHandler := handlers.NewCharacterHandler(s)

	character := s.Echo.Group("/characters")
	character.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	character.GET("", characterHandler.List)
}
