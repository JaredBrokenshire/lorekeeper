package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"lorekeeper/pkg/logging"
	"net/http"
	"strings"
)

type JWTClaims struct {
	UserID uint `json:"user_id"`

	jwt.RegisteredClaims
}

func JWTAuth(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Read the Authorization header.
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				logging.Logger.LogWarn().Msg("Missing or malformed jwt")
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Missing or malformed jwt",
				)
			}

			// Expect: Authorization: Bearer <token>
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				logging.Logger.LogWarn().Msg("Invalid authorization header")
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Invalid authorization header",
				)
			}

			tokenString := strings.TrimSpace(parts[1])
			if tokenString == "" {
				logging.Logger.LogWarn().Msg("Missing access token")
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Missing access token",
				)
			}

			claims := &JWTClaims{}

			token, err := jwt.ParseWithClaims(
				tokenString,
				claims,
				func(token *jwt.Token) (interface{}, error) {
					// Only accept the signing algorithm used by the backend.
					if token.Method != jwt.SigningMethodHS256 {
						logging.Logger.LogWarn().Msg("Unexpected signing method")
						return nil, echo.NewHTTPError(
							http.StatusUnauthorized,
							"Unexpected signing method",
						)
					}

					return []byte(secret), nil
				},
			)

			if err != nil || token == nil || !token.Valid {
				logging.Logger.LogWarn().Msg("Invalid or expired access token")
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Invalid or expired access token",
				)
			}

			if claims.UserID == 0 {
				logging.Logger.LogWarn().Msg("Invalid token claims")
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Invalid token claims",
				)
			}

			// Make the authenticated user's ID available to handlers.
			c.Set("userID", claims.UserID)

			return next(c)
		}
	}
}
