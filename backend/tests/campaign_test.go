package tests

import (
	"fmt"
	"lorekeeper/db/factories"
	"lorekeeper/db/models"
	"lorekeeper/tests/helpers"
	"net/http"
	"testing"
)

func TestCampaign_List(t *testing.T) {
	ts.ClearTable("campaigns")
	ts.ClearTable("characters")
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
