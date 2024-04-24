package model

import (
	"fmt"

	"gorm.io/gorm"
)

type DirectMessage struct {
	gorm.Model
	Name string
}

type DirectMessageMember struct {
	gorm.Model
	DirectMessageID int
	DirectMessage   DirectMessage
	UserID          int
	User            User
}

func (DirectMessage) IsNode()            {}
func (this DirectMessage) GetID() string { return fmt.Sprintf("dm:%d", this.Model.ID) }

type DirectMessagePost struct {
	gorm.Model
	DirectMessageID int
	DirectMessage   DirectMessage
	UserID          int
	User            User
	Content         string `json:"content"`
}

func (DirectMessagePost) IsNode()            {}
func (this DirectMessagePost) GetID() string { return fmt.Sprintf("dmp:%d", this.Model.ID) }
