package routes

import (
	"lorekeeper/api"
	"lorekeeper/api/handlers"
	mw "lorekeeper/pkg/middleware"
	"os"
)

func campaignFileRoutes(s *api.Server) {
	campaignFileHandler := handlers.NewCampaignFileHandler(s)

	campaignFile := s.Echo.Group("/campaigns/:id/files")
	campaignFile.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	campaignFile.POST("", campaignFileHandler.Create)
}
