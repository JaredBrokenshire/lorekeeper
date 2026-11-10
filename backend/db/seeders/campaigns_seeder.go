package seeders

import (
	"fmt"
	"lorekeeper/db/models"
	"lorekeeper/pkg/logging"
)

func (s *Seeder) SetCampaigns() {
	campaigns := []models.Campaign{
		{
			ID:              1,
			Name:            "Marblethorn",
			DungeonMasterID: 8, // sophie_bee
		},
		{
			ID:              2,
			Name:            "Dungeoneers",
			DungeonMasterID: 3, // Bowser3882
		},
	}

	for _, campaign := range campaigns {
		err := s.DB.Where("id = ?", campaign.ID).FirstOrCreate(&campaign).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating campaign with id `%v` in seeder: %v", campaign.ID, s.DB.Error))
		}
	}
}
