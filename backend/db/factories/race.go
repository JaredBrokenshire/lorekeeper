package factories

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"lorekeeper/db/models"
)

func NewRace(db *gorm.DB, race *models.Race) {
	fillRaceDefaults(race)
	err := db.Create(race).Error
	if err != nil {
		panic("error creating race in factory: " + err.Error())
	}
}

func fillRaceDefaults(race *models.Race) {
	if race.Name == "" {
		race.Name = random.String(16)
	}
}
