package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	UniversityName string `json:"university_name"`
	Name           string `json:"name"`
	Description    string `json:"description"`
}

func (Class) IsNode()            {}
func (this Class) GetID() string { return fmt.Sprintf("class:%d", this.Model.ID) }
