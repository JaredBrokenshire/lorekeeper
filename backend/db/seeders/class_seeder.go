package seeders

import (
	"fmt"
	"lorekeeper/db/models"
	"lorekeeper/pkg/logging"
)

func (s *Seeder) SetClasses() {
	classes := []models.Class{
		{
			ID:   1,
			Name: "Artificer",
		},
		{
			ID:   2,
			Name: "Barbarian",
		},
		{
			ID:   3,
			Name: "Bard",
		},
		{
			ID:   4,
			Name: "Cleric",
		},
		{
			ID:   5,
			Name: "Druid",
		},
		{
			ID:   6,
			Name: "Fighter",
		},
		{
			ID:   7,
			Name: "Monk",
		},
		{
			ID:   8,
			Name: "Paladin",
		},
		{
			ID:   9,
			Name: "Ranger",
		},
		{
			ID:   10,
			Name: "Rogue",
		},
		{
			ID:   11,
			Name: "Sorcerer",
		},
		{
			ID:   12,
			Name: "Warlock",
		},
		{
			ID:   13,
			Name: "Wizard",
		},
		{
			ID:   14,
			Name: "Blood Hunter",
		},
	}

	for _, c := range classes {
		err := s.DB.Where("id = ?", c.ID).FirstOrCreate(&c).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating class with id `%v` in seeder: %v", c.ID, err.Error()))
		}
	}
}
