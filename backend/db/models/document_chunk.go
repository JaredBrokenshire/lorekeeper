package models

import "time"

type DocumentChunk struct {
	ID uint `gorm:"primary_key;auto_increment"`

	CampaignFileID uint
	CampaignFile   *File `gorm:"foreignkey:CampaignFileID"`

	PageNumber int
	Index      int `gorm:"column:chunk_index"`
	Content    string

	CreatedAt time.Time
	UpdatedAt time.Time
}
