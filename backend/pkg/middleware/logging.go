package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"lorekeeper/pkg/logging"
	"net/http"
	"time"
)

func LoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			request := c.Request()

			// call the next middleware/handler
			err := next(c)

			status := c.Response().Status
			if err != nil {
				status = http.StatusInternalServerError
			}

			logFields := map[string]interface{}{
				"method":   request.Method,
				"status":   status,
				"duration": time.Since(start).String(),
			}

			uri := request.URL.Path
			if request.URL.RawQuery != "" {
				uri = fmt.Sprintf("%v?%v", request.URL.Path, request.URL.RawQuery)
			}

			logFields["uri"] = uri

			// log the request
			switch {
			case status >= 200 && status < 300:
				logging.Logger.WithLevel(zerolog.TraceLevel).Fields(logFields).Msg("Request")
				break
			case status >= 400 && status < 500:
				logging.Logger.LogWarn().Fields(logFields).Msg("Request")
				break
			case status >= 500:
				logging.Logger.LogError().Fields(logFields).Msg("Request")
				break
			default:
				logging.Logger.LogInfo().Fields(logFields).Msg("Request")
				break
			}

			return err
		}
	}
}
