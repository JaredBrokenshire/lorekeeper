package tests

import (
	"fmt"
	"lorekeeper/db/factories"
	"lorekeeper/db/models"
	"lorekeeper/tests/helpers"
	"net/http"
	"testing"
)

func TestCharacter_List(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("subclasses")
	ts.ClearTable("races")
	ts.ClearTable("characters")
	ts.ClearTable("character_classes")
	ts.ClearTable("character_subclasses")
	ts.ClearTable("campaigns")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	campaign := &models.Campaign{}
	factories.NewCampaign(ts.S.DB, campaign)

	race := &models.Race{}
	factories.NewRace(ts.S.DB, race)
	race2 := &models.Race{}
	factories.NewRace(ts.S.DB, race2)
	race3 := &models.Race{}
	factories.NewRace(ts.S.DB, race3)

	logo := &models.File{}
	factories.NewFile(ts.S.DB, logo)
	logo2 := &models.File{}
	factories.NewFile(ts.S.DB, logo2)
	logo3 := &models.File{}
	factories.NewFile(ts.S.DB, logo3)

	class := &models.Class{LogoID: logo.ID}
	factories.NewClass(ts.S.DB, class)
	class2 := &models.Class{LogoID: logo2.ID}
	factories.NewClass(ts.S.DB, class2)
	class3 := &models.Class{LogoID: logo3.ID}
	factories.NewClass(ts.S.DB, class3)

	subclass := &models.Subclass{ClassID: class.ID}
	factories.NewSubclass(ts.S.DB, subclass)
	subclass2 := &models.Subclass{ClassID: class2.ID}
	factories.NewSubclass(ts.S.DB, subclass2)
	subclass3 := &models.Subclass{ClassID: class3.ID}
	factories.NewSubclass(ts.S.DB, subclass3)

	profilePicture := &models.File{}
	factories.NewFile(ts.S.DB, profilePicture)
	profilePicture2 := &models.File{}
	factories.NewFile(ts.S.DB, profilePicture2)
	profilePicture3 := &models.File{}
	factories.NewFile(ts.S.DB, profilePicture3)
	profilePicture4 := &models.File{}
	factories.NewFile(ts.S.DB, profilePicture4)

	character := &models.Character{
		RaceID:           race.ID,
		CampaignID:       campaign.ID,
		Classes:          []*models.Class{class, class2},
		Subclasses:       []*models.Subclass{subclass, subclass2},
		ProfilePictureID: profilePicture.ID,
	}
	factories.NewCharacter(ts.S.DB, character)
	character2 := &models.Character{
		RaceID:           race2.ID,
		Classes:          []*models.Class{class2},
		Subclasses:       []*models.Subclass{subclass2},
		ProfilePictureID: profilePicture2.ID,
	}
	factories.NewCharacter(ts.S.DB, character2)
	namedCharacter := &models.Character{
		Name:             "Test Character",
		RaceID:           race3.ID,
		Classes:          []*models.Class{class3},
		Subclasses:       []*models.Subclass{subclass3},
		ProfilePictureID: profilePicture3.ID,
	}
	factories.NewCharacter(ts.S.DB, namedCharacter)
	otherUserCharacter := &models.Character{UserID: 999, ProfilePictureID: profilePicture4.ID}
	factories.NewCharacter(ts.S.DB, otherUserCharacter)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/characters%v", query),
		}
	}

	permissionRequest := getRequest("")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can get characters",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),

					fmt.Sprintf(`"level":%v`, character.Level),
					fmt.Sprintf(`"level":%v`, character2.Level),
					fmt.Sprintf(`"level":%v`, namedCharacter.Level),

					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, race3.Name),

					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, class3.Name),

					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),
					fmt.Sprintf(`"name":"%v"`, subclass3.Name),

					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"name":"%v"`, campaign.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture3.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture3.Location),

					fmt.Sprintf(`"name":"%v"`, logo.Name),
					fmt.Sprintf(`"location":"%v"`, logo.Location),
					fmt.Sprintf(`"name":"%v"`, logo2.Name),
					fmt.Sprintf(`"location":"%v"`, logo2.Location),
					fmt.Sprintf(`"name":"%v"`, logo3.Name),
					fmt.Sprintf(`"location":"%v"`, logo3.Location),

					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture4.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture4.Location),
				},
			},
		},
		{
			Name:    "Can get page 0 of characters",
			Request: getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"level":%v`, character.Level),

					fmt.Sprintf(`"name":"%v"`, race.Name),

					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, class2.Name),

					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),

					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"name":"%v"`, campaign.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),

					fmt.Sprintf(`"name":"%v"`, logo.Name),
					fmt.Sprintf(`"location":"%v"`, logo.Location),
					fmt.Sprintf(`"name":"%v"`, logo2.Name),
					fmt.Sprintf(`"location":"%v"`, logo2.Location),

					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),

					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, race3.Name),

					fmt.Sprintf(`"name":"%v"`, class3.Name),

					fmt.Sprintf(`"name":"%v"`, subclass3.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture3.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture3.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture4.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture4.Location),

					fmt.Sprintf(`"name":"%v"`, logo3.Name),
					fmt.Sprintf(`"location":"%v"`, logo3.Location),
				},
			},
		},
		{
			Name:    "Can get page 1 of characters",
			Request: getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"level":%v`, character2.Level),

					fmt.Sprintf(`"name":"%v"`, race2.Name),

					fmt.Sprintf(`"name":"%v"`, class2.Name),

					fmt.Sprintf(`"name":"%v"`, subclass2.Name),

					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),

					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),

					fmt.Sprintf(`"name":"%v"`, logo2.Name),
					fmt.Sprintf(`"location":"%v"`, logo2.Location),

					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),

					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, race3.Name),

					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, class3.Name),

					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"name":"%v"`, subclass3.Name),

					fmt.Sprintf(`"name":"%v"`, campaign.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture3.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture3.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture4.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture4.Location),

					fmt.Sprintf(`"name":"%v"`, logo.Name),
					fmt.Sprintf(`"location":"%v"`, logo.Location),
					fmt.Sprintf(`"name":"%v"`, logo3.Name),
					fmt.Sprintf(`"location":"%v"`, logo3.Location),
				},
			},
		},
		{
			Name:    "Can search for characters by name",
			Request: getRequest("?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"level":%v`, namedCharacter.Level),

					fmt.Sprintf(`"name":"%v"`, race3.Name),

					fmt.Sprintf(`"name":"%v"`, class3.Name),

					fmt.Sprintf(`"name":"%v"`, subclass3.Name),

					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),

					fmt.Sprintf(`"name":"%v"`, profilePicture3.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture3.Location),

					fmt.Sprintf(`"name":"%v"`, logo3.Name),
					fmt.Sprintf(`"location":"%v"`, logo3.Location),

					`"total_count":1`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),

					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),

					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, class2.Name),

					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),

					fmt.Sprintf(`"name":"%v"`, campaign.Name),

					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture4.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture4.Location),

					fmt.Sprintf(`"name":"%v"`, logo.Name),
					fmt.Sprintf(`"location":"%v"`, logo.Location),
					fmt.Sprintf(`"name":"%v"`, logo2.Name),
					fmt.Sprintf(`"location":"%v"`, logo2.Location),
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
