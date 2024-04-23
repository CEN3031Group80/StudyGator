package model

type DirectMessage struct {
	ID      string               `json:"id"`
	Members []*User              `json:"members"`
	Name    *string              `json:"name,omitempty"`
	Posts   []*DirectMessagePost `json:"posts"`
}

func (DirectMessage) IsNode()            {}
func (this DirectMessage) GetID() string { return this.ID }

type DirectMessagePost struct {
	ID            string         `json:"id"`
	DirectMessage *DirectMessage `json:"directMessage"`
	Sender        *User          `json:"sender"`
	Content       string         `json:"content"`
}

func (DirectMessagePost) IsNode()            {}
func (this DirectMessagePost) GetID() string { return this.ID }
