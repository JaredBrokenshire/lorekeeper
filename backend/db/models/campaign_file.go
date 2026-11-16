package models

type CampaignFile struct {
	ID uint `gorm:"primary_key;auto_increment"`

	CampaignID uint
	Campaign   *Campaign

	FileID uint
	File   *File

	Chunks []*DocumentChunk
}
