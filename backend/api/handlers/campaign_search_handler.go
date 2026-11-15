package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"lorekeeper/api"
	"lorekeeper/api/responses"
	"lorekeeper/pkg/logging"
	"lorekeeper/services"
)

type CampaignSearchHandler struct {
	server *api.Server
}

func NewCampaignSearchHandler(server *api.Server) *CampaignSearchHandler {
	return &CampaignSearchHandler{server: server}
}

// List godoc
// @Summary Search campaign session notes
// @Description Search the authenticated user's campaign session notes using BM25 relevance ranking.
// @ID campaigns-search-list
// @Tags Campaign Actions
// @Accept json
// @Produce json
// @Param id path int true "Campaign ID"
// @Param search query string true "Search query"
// @Param limit query int false "Maximum number of results (default 10, maximum 50)"
// @Success 200 {object} []responses.SearchResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /campaigns/:id/search [get]
func (h *CampaignSearchHandler) List(c echo.Context) error {
	userID := c.Get("userID").(uint)
	campaignID := c.Param("id")

	search := strings.TrimSpace(c.QueryParam("search"))
	if search == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, "A search query is required")
	}

	limit := services.DefaultSearchLimit
	if rawLimit := c.QueryParam("limit"); rawLimit != "" {
		parsedLimit, parseErr := strconv.Atoi(rawLimit)
		if parseErr != nil {
			return responses.ErrorResponse(c, http.StatusBadRequest, "Invalid search result limit")
		}
		limit = parsedLimit
	}

	results, err := h.server.Dependencies.GetSearchService().Search(campaignID, userID, search, limit)
	switch {
	case errors.Is(err, services.ErrSearchCampaignNotFound):
		return responses.ErrorResponse(c, http.StatusNotFound, "Campaign not found")
	case errors.Is(err, services.ErrSearchQueryTooLong):
		return responses.ErrorResponse(c, http.StatusBadRequest, "Search query is too long")
	case errors.Is(err, services.ErrSearchInvalidLimit):
		return responses.ErrorResponse(c, http.StatusBadRequest, "Invalid search result limit")
	case err != nil:
		logging.Logger.LogError().Msg(fmt.Sprintf("campaign/search error: %v", err))
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was an error searching campaign notes")
	}

	res := responses.NewSearchResponses(results)
	return responses.Response(c, http.StatusOK, res)
}
