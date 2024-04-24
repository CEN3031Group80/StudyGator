package model

import (
	"strconv"
	"strings"

	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	gorm_db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = gorm_db

	DB.AutoMigrate(&Class{})
	DB.AutoMigrate(&DirectMessage{})
	DB.AutoMigrate(&DirectMessageMember{})
	DB.AutoMigrate(&DirectMessagePost{})
	DB.AutoMigrate(&Post{})
	DB.AutoMigrate(&PostAttachment{})
	DB.AutoMigrate(&StudyGroup{})
	DB.AutoMigrate(&FriendRequest{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&StudyGroupMember{})

	DB.Create(&Class{
		UniversityName: "University of Florida",
		Name:           "CEN3031",
		Description:    "Intro to Software Engineering",
	})

	DB.Create(&Class{
		UniversityName: "University of Florida",
		Name:           "CDA4630",
		Description:    "Embedded Systems",
	})

	DB.Create(&Class{
		UniversityName: "University of Florida",
		Name:           "MAS3114",
		Description:    "Computational Linear Algebra",
	})

	return nil
}

func StringIDToIntID(id string) uint {
	parts := strings.Split(id, ":")
	if len(parts) < 2 {
		return 0
	}

	out, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return 0
	}

	return uint(out)
}
