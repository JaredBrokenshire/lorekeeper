package routes

import (
	"lorekeeper/api"
	"lorekeeper/api/handlers"
	mw "lorekeeper/pkg/middleware"
	"os"
)

func campaignRoutes(s *api.Server) {
	campaignHandler := handlers.NewCampaignHandler(s)

	campaign := s.Echo.Group("/campaigns")
	campaign.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	campaign.GET("", campaignHandler.List)
	campaign.GET("/:id", campaignHandler.Get)
}
