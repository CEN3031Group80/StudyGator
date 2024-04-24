package model

import (
	"fmt"

	"gorm.io/gorm"
)

type FriendRequest struct {
	gorm.Model
	SenderID   int
	Sender     User
	ReceiverID int
	Receiver   User
	Accepted   bool `json:"accepted"`
}

func (FriendRequest) IsNode()            {}
func (this FriendRequest) GetID() string { return fmt.Sprintf("fr:%d", this.Model.ID) }

type User struct {
	gorm.Model
	AltID          string        `gorm:"index:idx_alt_id,unique"`
	AvatarURL      string        `json:"avatarURL"`
	Provider       AuthProviders `json:"provider"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	FirstName      string        `json:"firstName"`
	LastName       string        `json:"lastName"`
	School         string        `json:"school"`
	GraduationYear int           `json:"graduationYear"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return fmt.Sprintf("user:%d", this.Model.ID) }
