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

type CampaignSearchHandler struct {
	server *api.Server
}

func NewCampaignSearchHandler(server *api.Server) *CampaignSearchHandler {
	return &CampaignSearchHandler{server: server}
}

// Ask godoc
// @Summary Ask a question about the content of campaign session notes
// @Description Returns an AI model response to a natural language question with citations to relevant campaign session note chunks, chosen through BM25 ranking
// @ID campaigns-search-ask
// @Tags Campaign Actions
// @Accept json
// @Produce json
// @Param id path int true "Campaign ID"
// @Param params body requests.CampaignSearchRequest true "Natural language question to rank session notes and prompt the AI model"
// @Success 200 {object} []responses.SearchResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /campaigns/:id/search [post]
func (h *CampaignSearchHandler) Ask(c echo.Context) error {
	userID := c.Get("userID").(uint)
	campaignID := c.Param("id")

	request := new(requests.CampaignSearchRequest)
	if err := c.Bind(request); err != nil {
		logging.Logger.LogError().Msg("campaigns/:id/search bind error: " + err.Error())
		return err
	}
	if err := c.Validate(request); err != nil {
		logging.Logger.LogError().Msg("campaigns/:id/search request validation error: " + err.Error())
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	questionResponse, err := h.server.Dependencies.GetQuestionService().Ask(campaignID, userID, c.Request().Context(), request.Question)
	switch {
	case errors.Is(err, services.ErrSearchRepository):
		logging.Logger.LogWarn().Msg("campaigns/:id/search search repository error" + err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was an reading from the database")
	case errors.Is(err, services.ErrCampaignNotFound):
		logging.Logger.LogWarn().Msg("campaigns/:id/search campaign not found" + err.Error())
		return responses.ErrorResponse(c, http.StatusNotFound, "Campaign not found")
	case errors.Is(err, services.ErrSearchNoResults):
		logging.Logger.LogInfo().Msg("campaigns/:id/search no search results")

		res := responses.NewQuestionResponse(nil)
		return responses.Response(c, http.StatusOK, res)
	case errors.Is(err, services.ErrOllamaPromptEmpty):
		logging.Logger.LogWarn().Msg("campaigns/:id/search ollama prompt empty" + err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Model prompt empty")
	case errors.Is(err, services.ErrOllamaContextEmpty):
		logging.Logger.LogWarn().Msg("campaigns/:id/search ollama context empty" + err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Model context empty")
	case errors.Is(err, services.ErrOllamaBuildPrompt):
		logging.Logger.LogWarn().Msg("campaigns/:id/search ollama build prompt error" + err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error building prompt")
	case errors.Is(err, services.ErrOllamaRequest):
		logging.Logger.LogWarn().Msg("campaigns/:id/search ollama request error" + err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error making request to model")
	case errors.Is(err, services.ErrOllamaResponse):
		logging.Logger.LogWarn().Msg("campaigns/:id/search ollama response error" + err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error parsing response from model")
	case errors.Is(err, services.ErrOllamaResponseEmpty):
		logging.Logger.LogWarn().Msg("campaigns/:id/search ollama response empty" + err.Error())
		return responses.ErrorResponse(c, http.StatusNotFound, "Model response not found")
	case err != nil:
		logging.Logger.LogError().Msg(fmt.Sprintf("campaigns/:id/search error: %v", err))
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was an error searching campaign notes")
	}

	res := responses.NewQuestionResponse(questionResponse)
	return responses.Response(c, http.StatusOK, res)
}
