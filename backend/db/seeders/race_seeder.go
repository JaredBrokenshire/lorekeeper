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
		{
			ID:   11,
			Name: "Dhampir",
		},
		{
			ID:   12,
			Name: "Warforged",
		},
		{
			ID:   13,
			Name: "Kobold",
		},
		{
			ID:   14,
			Name: "Half-Elf",
		},
		{
			ID:   15,
			Name: "Half-Orc",
		},
		{
			ID:   16,
			Name: "Air Genasi",
		},
		{
			ID:   17,
			Name: "Earth Genasi",
		},
		{
			ID:   18,
			Name: "Fire Genasi",
		},
		{
			ID:   19,
			Name: "Water Genasi",
		},
	}

	for _, r := range races {
		err := s.DB.Where("id = ?", r.ID).FirstOrCreate(&r).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating race with id `%v` in seeder: %v", r.ID, err.Error()))
		}
	}
}
