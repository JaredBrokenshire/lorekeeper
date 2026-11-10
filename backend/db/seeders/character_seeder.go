package seeders

import (
	"fmt"
	"lorekeeper/db/models"
	"lorekeeper/pkg/logging"
)

func (s *Seeder) SetCharacters() {
	characters := []struct {
		character   models.Character
		classIDs    []uint
		subclassIDs []uint
	}{
		{
			character: models.Character{
				ID:     1,
				UserID: 1, // JBrokenshire (admin)
				Name:   "Kael",
				Level:  7,
				RaceID: 2, // Dragonborn
			},
			classIDs:    []uint{6},   // Fighter
			subclassIDs: []uint{103}, // Echo Knight
		},
		{
			character: models.Character{
				ID:     2,
				UserID: 1, // JBrokenshire (admin)
				Name:   "Faelan",
				Level:  10,
				RaceID: 7, // Halfling
			},
			classIDs:    []uint{14}, // Blood Hunter
			subclassIDs: []uint{99}, // Order of the Mutant
		},
		{
			character: models.Character{
				ID:     3,
				UserID: 2, // TestUser
				Name:   "Ira",
				Level:  3,
				RaceID: 6, // Goliath
			},
			classIDs:    []uint{14},  // Blood Hunter
			subclassIDs: []uint{100}, // Order of the Lycan
		},
		{
			character: models.Character{
				ID:     4,
				UserID: 3, // Bowser3882
				Name:   "Niko",
				Level:  7,
				RaceID: 11, // Dhampir
			},
			classIDs:    []uint{10, 9}, // Rouge, Ranger
			subclassIDs: []uint{62},    // Gloom Stalker
		},
		{
			character: models.Character{
				ID:     5,
				UserID: 3, // Bowser3882
				Name:   "Blade",
				Level:  10,
				RaceID: 12, // Warforged
			},
			classIDs:    []uint{8},  // Paladin
			subclassIDs: []uint{54}, // Oath of the Ancients
		},
		{
			character: models.Character{
				ID:     6,
				UserID: 4, // TolandS
				Name:   "Idris",
				RaceID: 8, // Human
				Level:  7,
			},
			classIDs:    []uint{1}, // Artificer
			subclassIDs: []uint{4}, // Battle Smith
		},
		{
			character: models.Character{
				ID:     7,
				UserID: 4, // TolandS
				Name:   "Meddw",
				RaceID: 3, // Dwarf
				Level:  10,
			},
			classIDs:    []uint{7},   // Monk
			subclassIDs: []uint{104}, // Way of the Drunken Master
		},
		{
			character: models.Character{
				ID:     8,
				UserID: 5, // EM4L0013028
				Name:   "Karma",
				RaceID: 10, // Tiefling
				Level:  7,
			},
			classIDs:    []uint{10},  // Sorcerer
			subclassIDs: []uint{105}, // Shadow Magic
		},
		{
			character: models.Character{
				ID:     9,
				UserID: 6, // Babbles
				Name:   "Lycia",
				RaceID: 13, // Kobold
				Level:  6,
			},
			classIDs:    []uint{4},   // Cleric
			subclassIDs: []uint{106}, // Twilight Domain
		},
		{
			character: models.Character{
				ID:     10,
				UserID: 7, // Elfwnn
				Name:   "Vashta",
				RaceID: 15, // Half-Orc
				Level:  7,
			},
			classIDs:    []uint{2}, // Cleric
			subclassIDs: []uint{6}, // Path of the Berserker
		},
	}

	for _, data := range characters {
		character := data.character
		err := s.DB.Create(&character).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating character with id `%v` in seeder: %v", character.ID, err.Error()))
		}

		// Load character classes.
		var classes []*models.Class
		err = s.DB.Find(&classes, data.classIDs).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error loading classes for character with id `%v`: %v", character.ID, err.Error()))
			continue
		}

		// Associate classes.
		err = s.DB.Model(&character).Association("Classes").Replace(classes).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error associating classes for character with id `%v`: %v", character.ID, err.Error()))
		}

		// Load character subclasses.
		var subclasses []*models.Subclass
		err = s.DB.Find(&subclasses, data.subclassIDs).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error loading subclasses for character with id `%v`: %v", character.ID, err.Error()))
			continue
		}

		// Associate subclasses.
		err = s.DB.Model(&character).Association("Subclasses").Replace(subclasses).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error associating subclasses for character with id `%v`: %v", character.ID, err.Error()))
		}

	}
}
