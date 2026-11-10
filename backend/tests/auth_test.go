package tests

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"lorekeeper/api/requests"
	"lorekeeper/services"
	"lorekeeper/tests/helpers"
	"lorekeeper/tests/mocks"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAuth_Register(t *testing.T) {
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/register",
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't register without the required fields",
			Request:     request,
			RequestBody: requests.RegisterRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Username is a required field",
					"Email is a required field",
					"Password is a required field",
					"ConfirmPassword is a required field",
				},
			},
		},
		{
			Name:    "Can't register with an existing username",
			Request: request,
			RequestBody: requests.RegisterRequest{
				Username:        ts.TestUser.Username,
				Email:           "test-email@example.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Username is already taken",
			},
		},
		{
			Name:    "Can't register with an existing email",
			Request: request,
			RequestBody: requests.RegisterRequest{
				Username:        "test-username",
				Email:           ts.TestUser.Email,
				Password:        "password",
				ConfirmPassword: "password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Email is already registered",
			},
		},
		{
			Name:    "Can't register with a password that doesn't match the confirm password",
			Request: request,
			RequestBody: requests.RegisterRequest{
				Username:        "test-username",
				Email:           "test-email@example.com",
				Password:        "password",
				ConfirmPassword: "password2",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Passwords do not match",
			},
		},
		{
			Name: "Can't register if there is an error from the auth service",
			Setup: func(test *helpers.TestCase) {
				mockAuthService := mocks.NewMockAuthService()
				mockAuthService.RegisterError = errors.New("test error")

				ts.S.Dependencies.SetAuthService(mockAuthService)
			},
			Request: request,
			RequestBody: requests.RegisterRequest{
				Username:        "test-username",
				Email:           "test-email@example.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "There was an error registering",
			},
			Teardown: func() {
				ts.S.Dependencies.SetAuthService(
					services.NewAuthService(
						ts.S.DB,
						ts.S.Repos.User,
						ts.S.Repos.Role,
						os.Getenv("JWT_SECRET"),
					),
				)
			},
		},
		{
			Name:    "Can register with valid details",
			Request: request,
			RequestBody: requests.RegisterRequest{
				Username:        "test-username",
				Email:           "test-email@example.com",
				Password:        "test-password",
				ConfirmPassword: "test-password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"access_token":"`,
					`"username":"test-username"`,
					`"email":"test-email@example.com"`,
				},
				BodyPartsMissing: []string{
					"test-password",
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

func TestAuth_Login(t *testing.T) {
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/login",
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't login without the required fields",
			Request:     request,
			RequestBody: requests.LoginRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Username is a required field",
					"Password is a required field",
				},
			},
		},
		{
			Name:    "Can't login with username that doesn't exist",
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: "Invalid Username",
				Password: "Invalid Password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusUnauthorized,
				BodyPart:   "Invalid credentials",
			},
		},
		{
			Name:    "Can't login with incorrect password",
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: helpers.AdminUsername,
				Password: "Incorrect Password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusUnauthorized,
				BodyPart:   "Invalid credentials",
			},
		},
		{
			Name: "Can't login if there is an error from the auth service",
			Setup: func(test *helpers.TestCase) {
				mockAuthService := mocks.NewMockAuthService()
				mockAuthService.LoginError = errors.New("test error")

				ts.S.Dependencies.SetAuthService(mockAuthService)
			},
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: helpers.AdminUsername,
				Password: helpers.AdminPassword,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "There was an error logging in",
			},
			Teardown: func() {
				ts.S.Dependencies.SetAuthService(
					services.NewAuthService(
						ts.S.DB,
						ts.S.Repos.User,
						ts.S.Repos.Role,
						os.Getenv("JWT_SECRET"),
					),
				)
			},
		},
		{
			Name:    "Can login with correct credentials",
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: helpers.AdminUsername,
				Password: helpers.AdminPassword,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"access_token":"`,
					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
				},
				BodyPartsMissing: []string{
					helpers.AdminPassword,
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

func TestAuth_Refresh(t *testing.T) {
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/refresh",
	}

	// Obtain a valid refresh token by logging in first
	loginReq := helpers.Request{Method: http.MethodPost, Url: "/auth/login"}
	loginCase := helpers.TestCase{
		Request: loginReq,
		RequestBody: requests.LoginRequest{
			Username: helpers.AdminUsername,
			Password: helpers.AdminPassword,
		},
		Expected: helpers.ExpectedResponse{StatusCode: http.StatusOK},
	}
	loginRes := ts.ExecuteTestCase(&loginCase)
	var refreshCookie *http.Cookie
	for _, c := range loginRes.Result().Cookies() {
		if c.Name == "refresh_token" {
			refreshCookie = c
			break
		}
	}

	cases := []helpers.TestCase{
		{
			Name:    "Can't refresh without a refresh token cookie",
			Request: request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusUnauthorized,
				BodyPart:   "Refresh token not found",
			},
		},
		{
			Name:    "Can't refresh with an invalid refresh token",
			Request: request,
			RequestCookies: []*http.Cookie{
				{Name: "refresh_token", Value: "invalid.token.value"},
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusUnauthorized,
				BodyPart:   "Invalid refresh token",
			},
		},
		{
			Name: "Can't refresh when the user no longer exists",
			Setup: func(test *helpers.TestCase) {
				// Create a temporary user, get a refresh token, then delete the user
				tempUser := ts.AdminUser
				refreshToken, err := ts.S.Dependencies.GetAuthService().GenerateRefreshToken(tempUser)
				if err != nil {
					return
				}
				ts.ClearTable("users")
				test.RequestCookies = []*http.Cookie{
					{Name: "refresh_token", Value: refreshToken},
				}
			},
			Request: request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "User not found",
			},
			Teardown: func() {
				ts.SetupDefaultUsers()
				// Re-obtain the refresh cookie after users are recreated
			},
		},
		{
			Name: "Can't refresh if there is an error from the auth service",
			Setup: func(test *helpers.TestCase) {
				mockAuthService := mocks.NewMockAuthService()
				mockAuthService.RefreshError = errors.New("test error")

				ts.S.Dependencies.SetAuthService(mockAuthService)
			},
			Request: request,
			RequestCookies: []*http.Cookie{
				{Name: "refresh_token", Value: "any.token.value"},
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "There was a problem refreshing the token",
			},
			Teardown: func() {
				ts.S.Dependencies.SetAuthService(
					services.NewAuthService(
						ts.S.DB,
						ts.S.Repos.User,
						ts.S.Repos.Role,
						os.Getenv("JWT_SECRET"),
					),
				)
			},
		},
		{
			Name:    "Can refresh with a valid refresh token",
			Request: request,
			RequestCookies: []*http.Cookie{
				refreshCookie,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"access_token":"`,
					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunUnauthorisedTestCase(t, test)
		})
	}
}

func TestAuth_Me(t *testing.T) {
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/auth/me",
	}

	RunNoAuthenticationTests(t, request.Method, request.Url)

	cases := []helpers.TestCase{
		{
			Name: "Can't get current user if the user no longer exists",
			Setup: func(test *helpers.TestCase) {
				// Generate a valid access token for the admin user, then delete the user
				// so the handler receives a valid non-zero ID that doesn't exist in the DB.
				accessToken, err := ts.S.Dependencies.GetAuthService().GenerateAccessToken(ts.AdminUser)
				if err != nil {
					return
				}
				ts.ClearTable("users")
				test.RequestHeaders = map[string]string{
					"Authorization": "Bearer " + accessToken,
				}
			},
			Request: request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "User not found",
			},
			Teardown: func() {
				ts.SetupDefaultUsers()
			},
		},
		{
			Name:    "Can get current user when authenticated",
			Request: request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
					fmt.Sprintf(`"email":"%v"`, ts.AdminUser.Email),
					`"roles":`,
				},
				BodyPartsMissing: []string{
					helpers.AdminPassword,
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

func TestAuth_Logout(t *testing.T) {
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/logout",
	}

	RunNoAuthenticationTests(t, request.Method, request.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can logout",
			Request: request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Logged out successfully",
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					var foundCookie bool
					for _, c := range res.Result().Cookies() {
						if c.Name == "refresh_token" {
							foundCookie = true
							// Cookie should be cleared — value empty and MaxAge negative
							assert.Equal(t, "", c.Value)
							assert.Less(t, c.MaxAge, 0)
						}
					}
					assert.True(t, foundCookie, "refresh_token cookie should be present in response")
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
