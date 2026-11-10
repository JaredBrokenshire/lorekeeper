package factories

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"lorekeeper/db/models"
	"math/rand"
)

func NewCharacter(db *gorm.DB, character *models.Character) {
	fillCharacterDefaults(character)
	err := db.Create(character).Error
	if err != nil {
		panic("error creating character in factory: " + err.Error())
	}
}

func fillCharacterDefaults(character *models.Character) {
	if character.UserID == 0 {
		character.UserID = 1
	}
	if character.Name == "" {
		character.Name = random.String(16)
	}
	if character.Level == 0 {
		character.Level = rand.Intn(20)
	}
}
