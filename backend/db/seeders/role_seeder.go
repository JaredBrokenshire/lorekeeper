package seeders

import (
	"fmt"
	"lorekeeper/db/models"
	"lorekeeper/pkg/logging"
)

func (s *Seeder) SetRoles() {
	roles := []models.Role{
		{
			ID:   1,
			Name: models.RoleAdmin,
		},
		{
			ID:   2,
			Name: models.RoleUser,
		},
	}

	for _, role := range roles {
		err := s.DB.Where("id = ?", role.ID).FirstOrCreate(&role).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating role with id `%v` in seeder: %v", role.ID, s.DB.Error))
		}
	}
}
