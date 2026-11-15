package tests

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/random"
	"lorekeeper/db/factories"
	"lorekeeper/db/models"
	"lorekeeper/services"
	"lorekeeper/tests/helpers"
	"lorekeeper/tests/mocks"
	"net/http"
	"testing"
)

func TestCampaignSearch_List(t *testing.T) {
	ts.ClearTable("document_chunks")
	ts.ClearTable("campaign_files")
	ts.ClearTable("characters")
	ts.ClearTable("files")
	ts.ClearTable("campaigns")
	ts.SetupDefaultUsers()

	campaign := &models.Campaign{
		DungeonMasterID: ts.AdminUser.ID,
		Name:            "Brackenford",
	}
	factories.NewCampaign(ts.S.DB, campaign)
	privateCampaign := &models.Campaign{
		DungeonMasterID: ts.TestUser.ID,
		Name:            "Private campaign",
	}
	factories.NewCampaign(ts.S.DB, privateCampaign)

	character := &models.Character{
		UserID:     ts.TestUser.ID,
		CampaignID: campaign.ID,
	}
	factories.NewCharacter(ts.S.DB, character)

	file := &models.File{
		Name:     "brackenford-session.pdf",
		Location: "/campaigns/1/session-notes/brackenford-session.pdf",
		AuthorID: ts.AdminUser.ID,
	}
	factories.NewFile(ts.S.DB, file)
	privateFile := &models.File{
		Name:     "private-notes.pdf",
		Location: "/campaigns/private/private-notes.pdf",
		AuthorID: ts.TestUser.ID,
	}
	factories.NewFile(ts.S.DB, privateFile)

	campaignFile := &models.CampaignFile{
		CampaignID: campaign.ID,
		FileID:     file.ID,
	}
	factories.NewCampaignFile(ts.S.DB, campaignFile)
	privateCampaignFile := &models.CampaignFile{
		CampaignID: privateCampaign.ID,
		FileID:     privateFile.ID,
	}
	factories.NewCampaignFile(ts.S.DB, privateCampaignFile)

	relevantChunk := &models.DocumentChunk{
		CampaignFileID: campaignFile.ID,
		PageNumber:     2,
		Index:          0,
		Content:        "The drowned creatures emerged from the cellar beneath the old tower.",
	}
	factories.NewDocumentChunk(ts.S.DB, relevantChunk)
	unrelatedChunk := &models.DocumentChunk{
		CampaignFileID: campaignFile.ID,
		PageNumber:     1,
		Index:          1,
		Content:        "The party purchased supplies from a travelling merchant.",
	}
	factories.NewDocumentChunk(ts.S.DB, unrelatedChunk)
	privateChunk := &models.DocumentChunk{
		CampaignFileID: privateCampaignFile.ID,
		PageNumber:     1,
		Index:          0,
		Content:        "The drowned creatures guard the hidden treasure.",
	}
	factories.NewDocumentChunk(ts.S.DB, privateChunk)

	getRequest := func(id interface{}, query string) helpers.Request {
		fmt.Println(">>> ", query)
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/campaigns/%v/search%v", id, query),
		}
	}

	permissionRequest := getRequest(campaign.ID, "")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't search a campaign with an invalid campaign ID",
			Request: getRequest("invalid-id", "?search=test search"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Campaign not found",
			},
		},
		{
			Name:    "Can't search a campaign with an invalid search limit",
			Request: getRequest("invalid-id", "?search=test search&limit=invalid-limit"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Invalid search result limit",
			},
		},
		{
			Name:    "Can't search a campaign that doesn't exist",
			Request: getRequest(999, "?search=test search"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Campaign not found",
			},
		},
		{
			Name:    "Can't search a campaign that the user doesn't have access to",
			Request: getRequest(privateCampaign.ID, "?search=test search"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Campaign not found",
			},
		},
		{
			Name:    "Can't search a campaign with empty search query",
			Request: getRequest(campaign.ID, "?search="),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "A search query is required",
			},
		},
		{
			Name:    "Can't search a campaign with a search query that's too long",
			Request: getRequest(campaign.ID, fmt.Sprintf("?search=%v", random.String(services.MaxSearchQuerySize+1))),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Search query is too long",
			},
		},
		{
			Name:    "Can't search a campaign with a negative limit",
			Request: getRequest(campaign.ID, "?search=test&limit=-1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Invalid search result limit",
			},
		},
		{
			Name:    "Can't search a campaign with limit above the maximum",
			Request: getRequest(campaign.ID, fmt.Sprintf("?search=test&limit=%v", services.MaxSearchLimit+1)),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Invalid search result limit",
			},
		},
		{
			Name: "Can't search campaign if there is an error from the search service",
			Setup: func(test *helpers.TestCase) {
				mockSearchService := mocks.NewMockSearchService()
				mockSearchService.SearchError = errors.New("test error")

				ts.S.Dependencies.SetSearchService(mockSearchService)
			},
			Request: getRequest(campaign.ID, "?search=test search"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "There was an error searching campaign notes",
			},
			Teardown: func() {
				ts.S.Dependencies.SetSearchService(nil)
			},
		},
		{
			Name:    "Can search campaign for relevant results",
			Request: getRequest(campaign.ID, "?search=drowned+creatures"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"content":"%v"`, relevantChunk.Content),
					fmt.Sprintf(`"page_number":%v`, relevantChunk.PageNumber),
					fmt.Sprintf(`"chunk_index":%v`, relevantChunk.Index),
					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"name":"%v"`, file.Name),
					fmt.Sprintf(`"location":"%v"`, file.Location),
				},
				BodyPartsMissing: []string{
					"The party purchased supplies",
					"private-notes.pdf",
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
