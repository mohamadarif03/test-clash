package database

import (
	"log"
	"test-clash-be/internal/model"
	"test-clash-be/pkg/utils"
)

func SeedAdmin() {
	var count int64
	DB.Model(&model.User{}).Where("role = ?", model.RoleAdmin).Count(&count)

	if count == 0 {
		hashedPassword, err := utils.HashPassword("admin123") // Default password
		if err != nil {
			log.Fatalf("Gagal melakukan hash password admin: %v", err)
		}

		admin := model.User{
			Name:     "Admin Super",
			Email:    "admin@testclash.com",
			Password: hashedPassword,
			Role:     model.RoleAdmin,
			RankTier: model.RankDiamond,
		}

		if err := DB.Create(&admin).Error; err != nil {
			log.Printf("Gagal membuat user admin: %v", err)
		} else {
			log.Println("User admin berhasil dibuat (email: admin@testclash.com, password: admin123)")
		}
	} else {
		log.Println("User admin sudah ada, seeding dilewati.")
	}
}
