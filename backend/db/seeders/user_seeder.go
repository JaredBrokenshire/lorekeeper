package seeders

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"lorekeeper/db/models"
	"lorekeeper/pkg/logging"
	"os"
)

func (s *Seeder) SetUsers() {
	// Look up roles seeded by SetRoles so we associate by existing record
	// rather than letting GORM try to insert them again.
	var adminRole, userRole models.Role
	if err := s.DB.Where("name = ?", models.RoleAdmin).First(&adminRole).Error; err != nil {
		log.Printf("seeder: could not find admin role: %v", err)
		return
	}
	if err := s.DB.Where("name = ?", models.RoleUser).First(&userRole).Error; err != nil {
		log.Printf("seeder: could not find user role: %v", err)
		return
	}

	userPasswordHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	type seedUser struct {
		user models.User
		role models.Role
	}

	users := []seedUser{
		{
			user: models.User{
				ID:       1,
				Username: "JBrokenshire",
				Email:    os.Getenv("ADMIN_EMAIL"),
				Password: string(userPasswordHash),
			},
			role: adminRole,
		},
		{
			user: models.User{
				ID:       2,
				Username: "TestUser",
				Email:    "test@example.com",
				Password: string(userPasswordHash),
			},
			role: userRole,
		},
		{
			user: models.User{
				ID:       3,
				Username: "Bowser3882",
				Email:    "josh@email.com",
				Password: string(userPasswordHash),
			},
			role: userRole,
		},
		{
			user: models.User{
				ID:       4,
				Username: "TolandS",
				Email:    "sam@email.com",
				Password: string(userPasswordHash),
			},
			role: userRole,
		},
		{
			user: models.User{
				ID:       5,
				Username: "EM4L0013028",
				Email:    "em@email.com",
				Password: string(userPasswordHash),
			},
			role: userRole,
		},
		{
			user: models.User{
				ID:       6,
				Username: "Babbles",
				Email:    "abi@email.com",
				Password: string(userPasswordHash),
			},
			role: userRole,
		},
		{
			user: models.User{
				ID:       7,
				Username: "Elfwnn",
				Email:    "millie@email.com",
				Password: string(userPasswordHash),
			},
			role: userRole,
		},
		{
			user: models.User{
				ID:       8,
				Username: "sophie_bee",
				Email:    "sophie@email.com",
				Password: string(userPasswordHash),
			},
			role: userRole,
		},
	}

	for _, su := range users {
		existing := models.User{}
		if err := s.DB.Where("id = ?", su.user.ID).First(&existing).Error; err == nil {
			// User already exists — skip to keep seeder idempotent.
			continue
		}

		if err := s.DB.Create(&su.user).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating user `%v` in seeder: %v", su.user.Username, s.DB.Error))
			continue
		}

		// GORM v1: Append returns *gorm.Association; check s.DB.Error afterward.
		s.DB.Model(&su.user).Association("Roles").Append(su.role)
		if s.DB.Error != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error assigning role to user `%v` in seeder: %v", su.user.Username, s.DB.Error))
		}
	}
}
