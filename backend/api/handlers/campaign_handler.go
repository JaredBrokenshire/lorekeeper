package handlers

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"lorekeeper/api"
	"lorekeeper/api/responses"
	"lorekeeper/pkg/logging"
	"lorekeeper/services"
	"net/http"
	"strconv"
)

type CampaignHandler struct {
	server *api.Server
}

func NewCampaignHandler(server *api.Server) *CampaignHandler {
	return &CampaignHandler{server: server}
}

// List godoc
// @Summary List campaigns
// @Description Return a paginated list of campaigns where the user is either the owner of the campaign or has a character in it
// @ID campaigns-list
// @Tags Campaign Actions
// @Accept json
// @Produce json
// @Param search query string false "Search campaigns by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.PaginatedCampaignResponse
// @Router /campaigns [get]
func (h *CampaignHandler) List(c echo.Context) error {
	userID := c.Get("userID").(uint)

	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("campaigns.name LIKE ?", searchTerm)
		})
	}

	campaigns, page, pageSize := h.server.Dependencies.GetCampaignService().List(userID, page, pageSize, scopes)
	count := h.server.Dependencies.GetCampaignService().Count(userID, scopes)

	res := responses.NewPaginatedCampaignResponse(campaigns, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get campaign
// @Description Get a campaign by its ID where the user is either the owner of the campaign or has a character in it
// @ID campaigns-get
// @Tags Campaign Actions
// @Accept json
// @Produce json
// @Param id path int true "Campaign ID"
// @Success 200 {object} responses.CampaignResponse
// @Router /campaigns/:id [get]
func (h *CampaignHandler) Get(c echo.Context) error {
	userID := c.Get("userID").(uint)
	id := c.Param("id")

	campaign, err := h.server.Dependencies.GetCampaignService().Get(id, userID)
	switch {
	case errors.Is(err, services.ErrCampaignNotFound):
		logging.Logger.LogWarn().Msg("campaign/get campaign not found")
		return responses.ErrorResponse(c, http.StatusNotFound, "Campaign not found")
	case err != nil:
		logging.Logger.LogError().Msg(fmt.Sprintf("campaign/get error: %v", err.Error()))
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was an error getting the campaign")
	}

	res := responses.NewCampaignResponse(campaign)
	return responses.Response(c, http.StatusOK, res)
}
