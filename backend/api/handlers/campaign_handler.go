package handlers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"lorekeeper/api"
	"lorekeeper/api/responses"
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
