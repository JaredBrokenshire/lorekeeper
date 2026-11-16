package tests

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/random"
	"lorekeeper/api/requests"
	"lorekeeper/db/factories"
	"lorekeeper/db/models"
	"lorekeeper/services"
	"lorekeeper/tests/helpers"
	"lorekeeper/tests/mocks"
	"net/http"
	"testing"
)

func TestCampaignSearch_Ask(t *testing.T) {
	ts.ClearTable("campaigns")
	ts.ClearTable("campaign_files")
	ts.ClearTable("files")
	ts.ClearTable("document_chunks")
	ts.ClearTable("characters")
	ts.SetupDefaultUsers()

	mockOllamaService := mocks.NewMockOllamaService()

	setup := func(test *helpers.TestCase) {
		mockOllamaService.Reset()
		ts.S.Dependencies.SetOllamaService(mockOllamaService)

		ts.S.Dependencies.SetQuestionService(
			services.NewQuestionService(
				ts.S.Dependencies.GetSearchService(),
				ts.S.Dependencies.GetOllamaService(),
			),
		)
	}

	dmCampaign := &models.Campaign{}
	factories.NewCampaign(ts.S.DB, dmCampaign)
	playerCampaign := &models.Campaign{DungeonMasterID: 999}
	factories.NewCampaign(ts.S.DB, playerCampaign)
	neverReturnedCampaign := &models.Campaign{DungeonMasterID: 999}
	factories.NewCampaign(ts.S.DB, neverReturnedCampaign)
	noChunksCampaign := &models.Campaign{}
	factories.NewCampaign(ts.S.DB, noChunksCampaign)

	character := &models.Character{CampaignID: playerCampaign.ID}
	factories.NewCharacter(ts.S.DB, character)

	file := &models.File{}
	factories.NewFile(ts.S.DB, file)
	file2 := &models.File{}
	factories.NewFile(ts.S.DB, file2)
	file3 := &models.File{}
	factories.NewFile(ts.S.DB, file3)

	dmCampaignFile := &models.CampaignFile{CampaignID: dmCampaign.ID, FileID: file.ID}
	factories.NewCampaignFile(ts.S.DB, dmCampaignFile)
	playerCampaignFile := &models.CampaignFile{CampaignID: playerCampaign.ID, FileID: file2.ID}
	factories.NewCampaignFile(ts.S.DB, playerCampaignFile)
	neverReturnedCampaignFile := &models.CampaignFile{CampaignID: neverReturnedCampaign.ID, FileID: file3.ID}
	factories.NewCampaignFile(ts.S.DB, neverReturnedCampaignFile)

	chunk := &models.DocumentChunk{CampaignFileID: dmCampaignFile.ID, Content: "We fought some goblins in the forest"}
	factories.NewDocumentChunk(ts.S.DB, chunk)
	chunk2 := &models.DocumentChunk{CampaignFileID: dmCampaignFile.ID, Content: "There was a pirate ship in the harbor"}
	factories.NewDocumentChunk(ts.S.DB, chunk2)
	chunk3 := &models.DocumentChunk{CampaignFileID: playerCampaignFile.ID, Content: "There was an owlbear name Kittenwhiskers"}
	factories.NewDocumentChunk(ts.S.DB, chunk3)
	chunk4 := &models.DocumentChunk{CampaignFileID: playerCampaignFile.ID, Content: "The necromancers were chanting a ritual"}
	factories.NewDocumentChunk(ts.S.DB, chunk4)
	chunk5 := &models.DocumentChunk{CampaignFileID: playerCampaignFile.ID}
	factories.NewDocumentChunk(ts.S.DB, chunk5)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPost,
			Url:    fmt.Sprintf("/campaigns/%v/search", id),
		}
	}

	permissionRequest := getRequest(dmCampaign.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't ask question without required fields",
			Setup:   setup,
			Request: getRequest(dmCampaign.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid: ",
					"Question is a required field",
				},
			},
		},
		{
			Name:    "Can't ask a question with invalid request",
			Setup:   setup,
			Request: getRequest(dmCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: random.String(201),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid: ",
					"Question must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can't ask a question with invalid campaign id",
			Setup:   setup,
			Request: getRequest("invalid-id"),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Test question",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Campaign not found",
			},
		},
		{
			Name:    "Can't ask a question for campaign that doesn't exist",
			Setup:   setup,
			Request: getRequest(999),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Test question",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Campaign not found",
			},
		},
		{
			Name:    "Can't ask a question for campaign that the user doesn't have access to",
			Setup:   setup,
			Request: getRequest(neverReturnedCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Test question",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Campaign not found",
			},
		},
		{
			Name:    "Can return empty response if there are no chunks for a campaign",
			Setup:   setup,
			Request: getRequest(noChunksCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Test question",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   `"answer":"There was no relevant information found."`,
			},
		},
		{
			Name: "Can't ask question if there is an empty prompt error from the ollama service",
			Setup: func(test *helpers.TestCase) {
				setup(test)

				mockOllamaService.GenerateError = services.ErrOllamaPromptEmpty
				ts.S.Dependencies.SetOllamaService(mockOllamaService)
				ts.S.Dependencies.SetQuestionService(
					services.NewQuestionService(
						ts.S.Dependencies.GetSearchService(),
						ts.S.Dependencies.GetOllamaService(),
					),
				)
			},
			Request: getRequest(dmCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Goblins",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "Model prompt empty",
			},
		},
		{
			Name: "Can't ask question if there is an empty context error from the ollama service",
			Setup: func(test *helpers.TestCase) {
				setup(test)

				mockOllamaService.GenerateError = services.ErrOllamaContextEmpty
				ts.S.Dependencies.SetOllamaService(mockOllamaService)
				ts.S.Dependencies.SetQuestionService(
					services.NewQuestionService(
						ts.S.Dependencies.GetSearchService(),
						ts.S.Dependencies.GetOllamaService(),
					),
				)
			},
			Request: getRequest(dmCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Goblins",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "Model context empty",
			},
		},
		{
			Name: "Can't ask question if there is a prompt build error from the ollama service",
			Setup: func(test *helpers.TestCase) {
				setup(test)

				mockOllamaService.GenerateError = services.ErrOllamaBuildPrompt
				ts.S.Dependencies.SetOllamaService(mockOllamaService)
				ts.S.Dependencies.SetQuestionService(
					services.NewQuestionService(
						ts.S.Dependencies.GetSearchService(),
						ts.S.Dependencies.GetOllamaService(),
					),
				)
			},
			Request: getRequest(dmCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Goblins",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "Error building prompt",
			},
		},
		{
			Name: "Can't ask question if there is a model request error from the ollama service",
			Setup: func(test *helpers.TestCase) {
				setup(test)

				mockOllamaService.GenerateError = services.ErrOllamaRequest
				ts.S.Dependencies.SetOllamaService(mockOllamaService)
				ts.S.Dependencies.SetQuestionService(
					services.NewQuestionService(
						ts.S.Dependencies.GetSearchService(),
						ts.S.Dependencies.GetOllamaService(),
					),
				)
			},
			Request: getRequest(dmCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Goblins",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "Error making request to model",
			},
		},
		{
			Name: "Can't ask question if there is a model response error from the ollama service",
			Setup: func(test *helpers.TestCase) {
				setup(test)

				mockOllamaService.GenerateError = services.ErrOllamaResponse
				ts.S.Dependencies.SetOllamaService(mockOllamaService)
				ts.S.Dependencies.SetQuestionService(
					services.NewQuestionService(
						ts.S.Dependencies.GetSearchService(),
						ts.S.Dependencies.GetOllamaService(),
					),
				)
			},
			Request: getRequest(dmCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Goblins",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "Error parsing response from model",
			},
		},
		{
			Name: "Can't ask question if there is an empty model response error from the ollama service",
			Setup: func(test *helpers.TestCase) {
				setup(test)

				mockOllamaService.GenerateError = services.ErrOllamaResponseEmpty
				ts.S.Dependencies.SetOllamaService(mockOllamaService)
				ts.S.Dependencies.SetQuestionService(
					services.NewQuestionService(
						ts.S.Dependencies.GetSearchService(),
						ts.S.Dependencies.GetOllamaService(),
					),
				)
			},
			Request: getRequest(dmCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Goblins",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Model response not found",
			},
		},
		{
			Name: "Can't ask question if there is an unexpected error from the services",
			Setup: func(test *helpers.TestCase) {
				setup(test)

				mockOllamaService.GenerateError = errors.New("test error")
				ts.S.Dependencies.SetOllamaService(mockOllamaService)
				ts.S.Dependencies.SetQuestionService(
					services.NewQuestionService(
						ts.S.Dependencies.GetSearchService(),
						ts.S.Dependencies.GetOllamaService(),
					),
				)
			},
			Request: getRequest(dmCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Goblins",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "There was an error searching campaign notes",
			},
		},
		{
			Name:    "Can ask question for campaign where the user is the owner",
			Setup:   setup,
			Request: getRequest(dmCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Where were the goblins?",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"answer":"`,
					fmt.Sprintf(`"file_name":"%v"`, file.Name),
					fmt.Sprintf(`"page_number":%v`, chunk.PageNumber),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"page_number":%v`, chunk2.PageNumber),
					fmt.Sprintf(`"page_number":%v`, chunk3.PageNumber),
					fmt.Sprintf(`"page_number":%v`, chunk4.PageNumber),

					fmt.Sprintf(`"file_name":"%v"`, file2.Name),
					fmt.Sprintf(`"file_name":"%v"`, file3.Name),
				},
			},
		},
		{
			Name:    "Can ask question for campaign where the user is a player",
			Setup:   setup,
			Request: getRequest(playerCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "What is the name of the owlbear?",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"answer":"`,
					fmt.Sprintf(`"file_name":"%v"`, file2.Name),
					fmt.Sprintf(`"page_number":%v`, chunk3.PageNumber),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"page_number":%v`, chunk.PageNumber),
					fmt.Sprintf(`"page_number":%v`, chunk2.PageNumber),

					fmt.Sprintf(`"file_name":"%v"`, file.Name),
					fmt.Sprintf(`"file_name":"%v"`, file3.Name),
				},
			},
		},
		{
			Name:    "Can get empty response for question not covered by campaign notes",
			Setup:   setup,
			Request: getRequest(dmCampaign.ID),
			RequestBody: requests.CampaignSearchRequest{
				Question: "Who is the king of England?",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"answer":"There was no relevant information found."`,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}
