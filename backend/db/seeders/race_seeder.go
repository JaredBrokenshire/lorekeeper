package seeders

import (
	"fmt"
	"lorekeeper/db/models"
	"lorekeeper/pkg/logging"
)

func (s *Seeder) SetRaces() {
	races := []models.Race{
		{
			ID:   1,
			Name: "Aasimar",
		},
		{
			ID:   2,
			Name: "Dragonborn",
		},
		{
			ID:   3,
			Name: "Dwarf",
		},
		{
			ID:   4,
			Name: "Elf",
		},
		{
			ID:   5,
			Name: "Gnome",
		},
		{
			ID:   6,
			Name: "Goliath",
		},
		{
			ID:   7,
			Name: "Halfling",
		},
		{
			ID:   8,
			Name: "Human",
		},
		{
			ID:   9,
			Name: "Orc",
		},
		{
			ID:   10,
			Name: "Tiefling",
		},
	}

	for _, race := range races {
		err := s.DB.Create(&race).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating race with id `%v` in seeder: %v", race.ID, s.DB.Error))
		}
	}
}
