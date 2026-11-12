package handlers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"lorekeeper/api"
	"lorekeeper/api/requests"
	"lorekeeper/api/responses"
	"lorekeeper/pkg/logging"
	"lorekeeper/services"
	"net/http"
)

type CampaignFileHandler struct {
	server *api.Server
}

func NewCampaignFileHandler(server *api.Server) *CampaignFileHandler {
	return &CampaignFileHandler{server: server}
}

// Create godoc
// @Summary Create campaign file
// @Description Create a file associated with a campaign. The user must own the campaign or have a character in it.
// @ID campaigns-file-create
// @Tags Campaign File Actions
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Campaign ID"
// @Param file formData file true "File to upload"
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Success 201 {object} responses.CampaignResponse
// @Router /campaigns/:id/files [post]
func (h *CampaignFileHandler) Create(c echo.Context) error {
	userID := c.Get("userID").(uint)
	id := c.Param("id")

	file, err := requests.MultipartFileRequest(c.Request(), "file")
	if err != nil {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("campaign/file/create error parsing file from request: %v", err.Error()))
		return responses.ErrorResponse(c, http.StatusBadRequest, "Error parsing file from request")
	}

	fileDTO := services.UploadedFileDTO{
		Name:        file.Name,
		ContentType: file.ContentType,
		Size:        file.Size,
		Open:        file.Open,
	}

	campaign, err := h.server.Dependencies.GetCampaignService().UploadFile(id, userID, fileDTO)
	switch {
	case errors.Is(err, services.ErrCampaignNotFound):
		logging.Logger.LogWarn().Msg("campaign/file/create campaign not found")
		return responses.ErrorResponse(c, http.StatusNotFound, "Campaign not found")
	case errors.Is(err, services.ErrCampaignInvalidFile):
		logging.Logger.LogWarn().Msg(fmt.Sprintf("campaign/file/create invalid campaign file %v", err.Error()))
		return responses.ErrorResponse(c, http.StatusBadRequest, "Invalid file")
	case errors.Is(err, services.ErrCampaignFileUpload):
		logging.Logger.LogWarn().Msg(fmt.Sprintf("campaign/file/create error saving file to file store %v", err.Error()))
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error saving file")
	case errors.Is(err, services.ErrCampaignFileSave):
		logging.Logger.LogWarn().Msg(fmt.Sprintf("campaign/file/create error saving file record %v", err.Error()))
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error saving file record")
	case err != nil:
		logging.Logger.LogError().Msg(fmt.Sprintf("campaign/file/create error: %v", err.Error()))
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was an error uploading the file")
	}

	res := responses.NewCampaignResponse(campaign)
	return responses.Response(c, http.StatusCreated, res)
}
