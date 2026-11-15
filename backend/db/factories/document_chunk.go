package factories

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"lorekeeper/db/models"
	"math/rand"
)

func NewDocumentChunk(db *gorm.DB, chunk *models.DocumentChunk) {
	fillDocumentChunkDefaults(chunk)
	err := db.Create(chunk).Error
	if err != nil {
		panic("error creating document chunk in factory: " + err.Error())
	}
}

func fillDocumentChunkDefaults(chunk *models.DocumentChunk) {
	if chunk.PageNumber == 0 {
		chunk.PageNumber = rand.Intn(100)
	}
	if chunk.Content == "" {
		chunk.Content = random.String(32)
	}
	if chunk.Index == 0 {
		chunk.Index = rand.Intn(100)
	}
}
