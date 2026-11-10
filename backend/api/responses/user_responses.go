package responses

import "lorekeeper/db/models"

type UserResponse struct {
	*SlimUserResponse
	Email string `json:"email"`

	Roles []RoleResponse `json:"roles"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type SlimUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func NewUserResponse(user *models.User) *UserResponse {
	res := &UserResponse{
		SlimUserResponse: NewSlimUserResponse(user),
		Email:            user.Email,
	}

	if len(user.Roles) > 0 {
		res.Roles = NewRoleResponses(user.Roles)
	}

	return res
}

func NewSlimUserResponse(user *models.User) *SlimUserResponse {
	return &SlimUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
