package models

type RoleName string

const (
	RoleAdmin RoleName = "admin"
	RoleUser  RoleName = "user"
)

type Role struct {
	ID   uint     `gorm:"primary_key;auto_increment"`
	Name RoleName `gorm:"unique_index"`
}
