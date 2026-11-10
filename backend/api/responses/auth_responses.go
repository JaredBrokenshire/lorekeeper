package responses

import "lorekeeper/db/models"

type AuthResponse struct {
	AccessToken string        `json:"access_token"`
	User        *UserResponse `json:"user"`
}

func NewAuthResponse(user *models.User, accessToken string) *AuthResponse {
	return &AuthResponse{
		AccessToken: accessToken,
		User:        NewUserResponse(user),
	}
}
