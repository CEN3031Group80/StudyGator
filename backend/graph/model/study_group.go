package model

type StudyGroup struct {
	ID          string  `json:"id"`
	Owner       *User   `json:"owner"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	JoinID      string  `json:"joinID"`
	Class       *Class  `json:"class"`
	Posts       []*Post `json:"posts"`
}

func (StudyGroup) IsNode()            {}
func (this StudyGroup) GetID() string { return this.ID }
