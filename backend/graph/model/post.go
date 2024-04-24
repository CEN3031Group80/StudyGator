package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	StudyGroupID int
	StudyGroup   StudyGroup
	PosterID     int
	Poster       User
	Name         string `json:"name"`
	Content      string `json:"content"`
}

func (Post) IsNode()            {}
func (this Post) GetID() string { return fmt.Sprintf("post:%d", this.Model.ID) }

type PostAttachment struct {
	gorm.Model
	UploadUUID  string `json:"uploadUUID"`
	Description string `json:"description,omitempty"`
	FileName    string `json:"fileName"`
}

func (PostAttachment) IsNode()            {}
func (this PostAttachment) GetID() string { return fmt.Sprintf("posta:%d", this.Model.ID) }
