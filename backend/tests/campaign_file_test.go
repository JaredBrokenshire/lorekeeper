package tests

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"lorekeeper/db/factories"
	"lorekeeper/db/models"
	"lorekeeper/services"
	"lorekeeper/tests/helpers"
	"lorekeeper/tests/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCampaignFile_Create(t *testing.T) {
	ts.ClearTable("campaigns")
	ts.ClearTable("files")
	ts.ClearTable("campaign_files")
	ts.SetupDefaultUsers()

	mockFileService := mocks.NewMockFileService()
	mockFileStore := mocks.NewMockFileStore()

	setup := func(test *helpers.TestCase) {
		ts.ClearTable("files")
		mockFileService.Reset()
		mockFileStore.Reset()

		ts.S.Dependencies.SetFileStore(mockFileStore)
	}

	dm := &models.User{}
	factories.NewUser(ts.S.DB, dm)

	dmCampaign := &models.Campaign{}
	factories.NewCampaign(ts.S.DB, dmCampaign)
	playerCampaign := &models.Campaign{DungeonMasterID: dm.ID}
	factories.NewCampaign(ts.S.DB, playerCampaign)
	noPermissionCampaign := &models.Campaign{DungeonMasterID: 999}
	factories.NewCampaign(ts.S.DB, noPermissionCampaign)

	character := &models.Character{CampaignID: playerCampaign.ID}
	factories.NewCharacter(ts.S.DB, character)

	pdf, pdfMw := createMultipartFile(t, "file", "../assets/test.pdf")
	pdf2, pdfMw2 := createMultipartFile(t, "file", "../assets/test.pdf")
	pdf3, pdfMw3 := createMultipartFile(t, "file", "../assets/test.pdf")
	pdf4, pdfMw4 := createMultipartFile(t, "file", "../assets/test.pdf")
	pdf5, pdfMw5 := createMultipartFile(t, "file", "../assets/test.pdf")
	pdf6, pdfMw6 := createMultipartFile(t, "file", "../assets/test.pdf")
	pdf7, pdfMw7 := createMultipartFile(t, "file", "../assets/test.pdf")
	pdf8, pdfMw8 := createMultipartFile(t, "file", "../assets/test.pdf")

	jpeg, jpegMw := createMultipartFile(t, "file", "../assets/test.jpeg")

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPost,
			Url:    fmt.Sprintf("/campaigns/%v/files", id),
		}
	}

	permissionRequest := getRequest(dmCampaign.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't create campaign file with no file",
			Setup:   setup,
			Request: getRequest(dmCampaign.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Error parsing file from request",
			},
		},
		{
			Name:               "Can't create campaign file with invalid id",
			Setup:              setup,
			Request:            getRequest("invalid-id"),
			RequestReader:      pdf,
			RequestContentType: pdfMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Campaign not found",
			},
		},
		{
			Name:               "Can't create campaign file for campaign that doesn't exist",
			Setup:              setup,
			Request:            getRequest(999),
			RequestReader:      pdf2,
			RequestContentType: pdfMw2.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Campaign not found",
			},
		},
		{
			Name:               "Can't create campaign file if user doesn't have permission",
			Setup:              setup,
			Request:            getRequest(noPermissionCampaign.ID),
			RequestReader:      pdf3,
			RequestContentType: pdfMw3.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Campaign not found",
			},
		},
		{
			Name: "Can't create campaign file if there is an error from file service Validate()",
			Setup: func(test *helpers.TestCase) {
				setup(test)

				mockFileService.ValidateError = errors.New("test error")
				ts.S.Dependencies.SetFileService(mockFileService)
				ts.S.Dependencies.SetCampaignService(services.NewCampaignService(
					ts.S.Repos.Campaign,
					ts.S.Dependencies.GetFileService(),
					ts.S.Dependencies.GetFileStore(),
				))
			},
			Request:            getRequest(dmCampaign.ID),
			RequestReader:      pdf4,
			RequestContentType: pdfMw4.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Invalid file",
			},
			Teardown: func() {
				ts.S.Dependencies.SetFileService(nil)
			},
		},
		{
			Name: "Can't create campaign file if there is an error from file store Save()",
			Setup: func(test *helpers.TestCase) {
				setup(test)

				mockFileStore.SaveError = errors.New("test error")
				ts.S.Dependencies.SetFileStore(mockFileStore)
				ts.S.Dependencies.SetCampaignService(services.NewCampaignService(
					ts.S.Repos.Campaign,
					ts.S.Dependencies.GetFileService(),
					ts.S.Dependencies.GetFileStore(),
				))
			},
			Request:            getRequest(dmCampaign.ID),
			RequestReader:      pdf5,
			RequestContentType: pdfMw5.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "Error saving file",
			},
			Teardown: func() {
				mockFileStore.Reset()
			},
		},
		{
			Name: "Can't create campaign file if there is an error from file service Create()",
			Setup: func(test *helpers.TestCase) {
				setup(test)

				mockFileService.CreateError = errors.New("test error")
				ts.S.Dependencies.SetFileService(mockFileService)
				ts.S.Dependencies.SetCampaignService(services.NewCampaignService(
					ts.S.Repos.Campaign,
					ts.S.Dependencies.GetFileService(),
					ts.S.Dependencies.GetFileStore(),
				))
			},
			Request:            getRequest(dmCampaign.ID),
			RequestReader:      pdf6,
			RequestContentType: pdfMw6.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "Error saving file record",
			},
			Teardown: func() {
				ts.S.Dependencies.SetFileService(nil)
			},
		},
		{
			Name: "Can't upload file with the wrong type",
			Setup: func(test *helpers.TestCase) {
				setup(test)
				ts.S.Dependencies.SetCampaignService(
					services.NewCampaignService(
						ts.S.Repos.Campaign,
						ts.S.Dependencies.GetFileService(),
						ts.S.Dependencies.GetFileStore(),
					),
				)
			},
			Request:            getRequest(dmCampaign.ID),
			RequestReader:      jpeg,
			RequestContentType: jpegMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Invalid file",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name:          "File record not created",
					Model:         models.File{},
					CountExpected: 0,
				},
			},
		},
		{
			Name: "Can upload valid PDF as a DM",
			Setup: func(test *helpers.TestCase) {
				setup(test)
				ts.S.Dependencies.SetFileStore(mockFileStore)
				ts.S.Dependencies.SetCampaignService(services.NewCampaignService(
					ts.S.Repos.Campaign,
					ts.S.Dependencies.GetFileService(),
					ts.S.Dependencies.GetFileStore(),
				))
			},
			Request:            getRequest(dmCampaign.ID),
			RequestReader:      pdf7,
			RequestContentType: pdfMw7.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, dmCampaign.Name),
					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"location":"/campaigns/%v/session-notes/`, dmCampaign.ID),
					".pdf",
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name:          "File record created",
					Model:         models.File{},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					assert.Equal(t, 1, len(mockFileStore.SaveCalls))
					assert.Contains(t, mockFileStore.SaveCalls[0].Path, fmt.Sprintf("/campaigns/%v/session-notes/", dmCampaign.ID))
					assert.Contains(t, mockFileStore.SaveCalls[0].FileName, ".pdf")
				},
			},
		},
		{
			Name: "Can upload valid PDF as a player",
			Setup: func(test *helpers.TestCase) {
				setup(test)
				ts.S.Dependencies.SetFileStore(mockFileStore)
				ts.S.Dependencies.SetCampaignService(services.NewCampaignService(
					ts.S.Repos.Campaign,
					ts.S.Dependencies.GetFileService(),
					ts.S.Dependencies.GetFileStore(),
				))
			},
			Request:            getRequest(playerCampaign.ID),
			RequestReader:      pdf8,
			RequestContentType: pdfMw8.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, playerCampaign.Name),
					fmt.Sprintf(`"username":"%v"`, dm.Username),
					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"location":"/campaigns/%v/session-notes/`, playerCampaign.ID),
					".pdf",
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name:          "File record created",
					Model:         models.File{},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					assert.Equal(t, 1, len(mockFileStore.SaveCalls))
					assert.Contains(t, mockFileStore.SaveCalls[0].Path, fmt.Sprintf("/campaigns/%v/session-notes/", playerCampaign.ID))
					assert.Contains(t, mockFileStore.SaveCalls[0].FileName, ".pdf")
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
