package factories

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"lorekeeper/db/models"
)

func NewUser(db *gorm.DB, user *models.User) {
	fillUserDefaults(user)
	err := db.Create(user).Error
	if err != nil {
		panic("error creating user in factory: " + err.Error())
	}
}

func fillUserDefaults(user *models.User) {
	if user.Username == "" {
		user.Username = random.String(16)
	}
	if user.Password == "" {
		user.Password = random.String(16)
	}
	if user.Email == "" {
		user.Email = random.String(16) + "@example.com"
	}
}
