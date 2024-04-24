package model

import (
	"fmt"

	"gorm.io/gorm"
)

type StudyGroup struct {
	gorm.Model
	OwnerID     int
	Owner       User
	Name        string `json:"name"`
	Description string `json:"description"`
	JoinID      string `json:"joinID"`
	ClassID     int
	Class       Class
}

func (StudyGroup) IsNode()            {}
func (this StudyGroup) GetID() string { return fmt.Sprintf("user:%d", this.Model.ID) }

type StudyGroupMember struct {
	gorm.Model
	Favorite     bool
	StudyGroupID int
	StudyGroup   StudyGroup
	UserID       int
	User         User
}
