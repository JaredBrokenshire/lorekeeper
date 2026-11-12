package tests

import (
	"errors"
	"fmt"
	"lorekeeper/db/factories"
	"lorekeeper/db/models"
	"lorekeeper/tests/helpers"
	"lorekeeper/tests/mocks"
	"net/http"
	"testing"
)

func TestCampaign_List(t *testing.T) {
	ts.ClearTable("campaigns")
	ts.ClearTable("characters")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	campaign := &models.Campaign{}
	factories.NewCampaign(ts.S.DB, campaign)
	otherUserCampaign := &models.Campaign{DungeonMasterID: ts.TestUser.ID}
	factories.NewCampaign(ts.S.DB, otherUserCampaign)
	namedCampaign := &models.Campaign{Name: "Test Campaign"}
	factories.NewCampaign(ts.S.DB, namedCampaign)
	neverReturnCampaign := &models.Campaign{DungeonMasterID: 999}
	factories.NewCampaign(ts.S.DB, neverReturnCampaign)

	profilePicture := &models.File{}
	factories.NewFile(ts.S.DB, profilePicture)
	profilePicture2 := &models.File{}
	factories.NewFile(ts.S.DB, profilePicture2)
	profilePicture3 := &models.File{}
	factories.NewFile(ts.S.DB, profilePicture3)

	character := &models.Character{CampaignID: otherUserCampaign.ID, ProfilePictureID: profilePicture.ID}
	factories.NewCharacter(ts.S.DB, character)
	otherUserCharacter := &models.Character{UserID: ts.TestUser.ID, CampaignID: campaign.ID, ProfilePictureID: profilePicture2.ID}
	factories.NewCharacter(ts.S.DB, otherUserCharacter)
	neverReturnCharacter := &models.Character{UserID: 999, CampaignID: neverReturnCampaign.ID, ProfilePictureID: profilePicture3.ID}
	factories.NewCharacter(ts.S.DB, neverReturnCharacter)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/campaigns%v", query),
		}
	}

	permissionRequest := getRequest("")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can get campaigns",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, namedCampaign.Name),

					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"username":"%v"`, ts.TestUser.Username),

					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),

					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, neverReturnCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnCharacter.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture3.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture3.Location),
				},
			},
		},
		{
			Name:    "Can get page 0 of campaigns",
			Request: getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, campaign.Name),

					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"username":"%v"`, ts.TestUser.Username),

					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),

					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, otherUserCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, namedCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnCharacter.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture3.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture3.Location),
				},
			},
		},
		{
			Name:    "Can get page 1 of campaigns",
			Request: getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, otherUserCampaign.Name),

					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"username":"%v"`, ts.TestUser.Username),

					fmt.Sprintf(`"name":"%v"`, character.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),

					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, namedCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnCharacter.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture3.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture3.Location),
				},
			},
		},
		{
			Name:    "Can search for campaigns by name",
			Request: getRequest("?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedCampaign.Name),

					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnCharacter.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture3.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture3.Location),
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

func TestCampaign_Get(t *testing.T) {
	ts.ClearTable("races")
	ts.ClearTable("classes")
	ts.ClearTable("subclasses")
	ts.ClearTable("campaigns")
	ts.ClearTable("campaign_files")
	ts.ClearTable("characters")
	ts.ClearTable("character_classes")
	ts.ClearTable("character_subclasses")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	author := &models.User{}
	factories.NewUser(ts.S.DB, author)

	sessionNote := &models.File{AuthorID: author.ID}
	factories.NewFile(ts.S.DB, sessionNote)
	sessionNote2 := &models.File{AuthorID: author.ID}
	factories.NewFile(ts.S.DB, sessionNote2)

	dmCampaign := &models.Campaign{DungeonMasterID: ts.AdminUser.ID, Files: []*models.File{sessionNote}}
	factories.NewCampaign(ts.S.DB, dmCampaign)
	playerCampaign := &models.Campaign{DungeonMasterID: ts.TestUser.ID, Files: []*models.File{sessionNote2}}
	factories.NewCampaign(ts.S.DB, playerCampaign)
	notReturnedCampaign := &models.Campaign{DungeonMasterID: 999}
	factories.NewCampaign(ts.S.DB, notReturnedCampaign)

	profilePicture := &models.File{}
	factories.NewFile(ts.S.DB, profilePicture)

	logo := &models.File{}
	factories.NewFile(ts.S.DB, logo)

	class := &models.Class{LogoID: logo.ID}
	factories.NewClass(ts.S.DB, class)

	subclass := &models.Subclass{ClassID: class.ID}
	factories.NewSubclass(ts.S.DB, subclass)

	race := &models.Race{}
	factories.NewRace(ts.S.DB, race)

	character := &models.Character{
		RaceID:           race.ID,
		Classes:          []*models.Class{class},
		Subclasses:       []*models.Subclass{subclass},
		CampaignID:       playerCampaign.ID,
		ProfilePictureID: profilePicture.ID,
	}
	factories.NewCharacter(ts.S.DB, character)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/campaigns/%v", id),
		}
	}

	permissionRequest := getRequest(1)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't get campaign with invalid ID",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyParts: []string{
					"Campaign not found",
				},
			},
		},
		{
			Name:    "Can't get campaign that doesn't exist",
			Request: getRequest(999),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyParts: []string{
					"Campaign not found",
				},
			},
		},
		{
			Name:    "Can't get campaign that doesn't belong to the user",
			Request: getRequest(notReturnedCampaign.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyParts: []string{
					"Campaign not found",
				},
			},
		},
		{
			Name:    "Can't get campaign if there is an error from the campaign service",
			Request: getRequest(dmCampaign.ID),
			Setup: func(test *helpers.TestCase) {
				mockCampaignService := mocks.NewMockCampaignService()
				mockCampaignService.GetError = errors.New("test error")

				ts.S.Dependencies.SetCampaignService(mockCampaignService)
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyParts: []string{
					"There was an error getting the campaign",
				},
			},
			Teardown: func() {
				ts.S.Dependencies.SetCampaignService(nil)
			},
		},
		{
			Name:    "Can get campaign where the user is the dungeon master",
			Request: getRequest(dmCampaign.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, dmCampaign.Name),

					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"username":"%v"`, author.Username),

					fmt.Sprintf(`"name":"%v"`, sessionNote.Name),
					fmt.Sprintf(`"location":"%v"`, sessionNote.Location),

					`"characters":null`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"username":"%v"`, ts.TestUser.Username),

					fmt.Sprintf(`"name":"%v"`, playerCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, notReturnedCampaign.Name),

					fmt.Sprintf(`"name":"%v"`, character.Name),

					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, subclass.Name),

					fmt.Sprintf(`"name":"%v"`, logo.Name),
					fmt.Sprintf(`"location":"%v"`, logo.Location),

					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),

					fmt.Sprintf(`"name":"%v"`, sessionNote2.Name),
					fmt.Sprintf(`"location":"%v"`, sessionNote2.Location),
				},
			},
		},
		{
			Name:    "Can get campaign where the user is a player",
			Request: getRequest(playerCampaign.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, playerCampaign.Name),

					fmt.Sprintf(`"username":"%v"`, ts.TestUser.Username),
					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"username":"%v"`, author.Username),

					fmt.Sprintf(`"name":"%v"`, character.Name),

					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, subclass.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),

					fmt.Sprintf(`"name":"%v"`, sessionNote2.Name),
					fmt.Sprintf(`"location":"%v"`, sessionNote2.Location),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, dmCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, notReturnedCampaign.Name),

					fmt.Sprintf(`"name":"%v"`, sessionNote.Name),
					fmt.Sprintf(`"location":"%v"`, sessionNote.Location),
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
