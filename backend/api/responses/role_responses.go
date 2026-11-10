package responses

import "lorekeeper/db/models"

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewRoleResponse(role *models.Role) *RoleResponse {
	return &RoleResponse{
		ID:   role.ID,
		Name: string(role.Name),
	}
}

func NewRoleResponses(roles []*models.Role) []RoleResponse {
	roleResponses := make([]RoleResponse, len(roles))
	for i, role := range roles {
		roleResponses[i] = *NewRoleResponse(role)
	}
	return roleResponses
}
