package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"lorekeeper/api"
	mw "lorekeeper/pkg/middleware"
	"os"
	"strings"
)

func RegisterRoutes(s *api.Server) {
	s.Echo.Use(mw.LoggingMiddleware())
	s.Echo.Use(middleware.Recover())

	// CORS origins are configured via the CORS_ORIGINS env var (comma-separated).
	// Falls back to the development default so existing local setups keep working.
	rawOrigins := os.Getenv("CORS_ORIGINS")
	if rawOrigins == "" {
		rawOrigins = "http://jb-lorekeeper.com:8080"
	}
	allowedOrigins := strings.Split(rawOrigins, ",")
	for i, o := range allowedOrigins {
		allowedOrigins[i] = strings.TrimSpace(o)
	}

	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		AllowCredentials: true,
	}))

	authRoutes(s)
	fileRoutes(s)
	characterRoutes(s)
	campaignRoutes(s)
}
