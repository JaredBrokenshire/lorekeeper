package factories

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"lorekeeper/db/models"
)

func NewClass(db *gorm.DB, class *models.Class) {
	fillClassDefaults(class)
	err := db.Create(class).Error
	if err != nil {
		panic("error creating class in factory: " + err.Error())
	}
}

func fillClassDefaults(class *models.Class) {
	if class.Name == "" {
		class.Name = random.String(16)
	}
}
