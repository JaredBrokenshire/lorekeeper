package routes

import (
	"github.com/labstack/echo/v4/middleware"
	"lorekeeper/api"
	"lorekeeper/api/handlers"
	"lorekeeper/config"
	mw "lorekeeper/pkg/middleware"
	"os"
)

func authRoutes(s *api.Server) {
	authHandler := handlers.NewAuthHandler(s)

	auth := s.Echo.Group("/auth")
	limiterStore := middleware.NewRateLimiterMemoryStore(config.Get().LoginRateLimit)

	auth.POST("/register", authHandler.Register, middleware.RateLimiter(limiterStore))
	auth.POST("/login", authHandler.Login, middleware.RateLimiter(limiterStore))
	auth.POST("/refresh", authHandler.Refresh, middleware.RateLimiter(limiterStore))

	protected := auth.Group("")
	protected.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	protected.GET("/me", authHandler.Me)
	protected.POST("/logout", authHandler.Logout)
}
