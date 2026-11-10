package handlers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"lorekeeper/api"
	"lorekeeper/api/responses"
	"lorekeeper/pkg/logging"
	"net/http"
	"strconv"
)

type CharacterHandler struct {
	server *api.Server
}

func NewCharacterHandler(server *api.Server) *CharacterHandler {
	return &CharacterHandler{server: server}
}

// List godoc
// @Summary List characters
// @Description Return a paginated list of characters belonging to the user
// @ID characters-list
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param search query string false "Search characters by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.PaginatedCharacterResponse
// @Router /characters [get]
func (h *CharacterHandler) List(c echo.Context) error {
	userID := c.Get("userID").(uint)
	if userID == 0 {
		logging.Logger.LogWarn().Msg("auth/me user not found")
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", searchTerm)
		})
	}

	characters, page, pageSize := h.server.Dependencies.GetCharacterService().List(userID, page, pageSize, scopes)
	count := h.server.Dependencies.GetCharacterService().Count(userID, scopes)

	res := responses.NewPaginatedCharacterResponse(characters, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}
