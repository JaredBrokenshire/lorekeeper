package factories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"lorekeeper/db/models"
)

func NewFile(db *gorm.DB, file *models.File) {
	fillFileDefaults(file)
	err := db.Create(file).Error
	if err != nil {
		panic("error creating file in factory: " + err.Error())
	}
}

func fillFileDefaults(file *models.File) {
	if file.Name == "" {
		file.Name = fmt.Sprintf("%v.png", random.String(16))
	}
	if file.Location == "" {
		file.Location = fmt.Sprintf("%v/%v/%v.png", random.String(16), random.String(16), random.String(16))
	}
	if file.AuthorID == 0 {
		file.AuthorID = 1
	}
}
