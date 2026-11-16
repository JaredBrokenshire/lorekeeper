package routes

import (
	"lorekeeper/api"
	"lorekeeper/api/handlers"
	mw "lorekeeper/pkg/middleware"
	"os"
)

func campaignSearchRoutes(s *api.Server) {
	campaignSearchHandler := handlers.NewCampaignSearchHandler(s)

	campaignSearch := s.Echo.Group("/campaigns/:id/search")
	campaignSearch.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	campaignSearch.POST("", campaignSearchHandler.Ask)
}
