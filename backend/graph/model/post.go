package model

type Post struct {
	ID          string            `json:"id"`
	StudyGroup  *StudyGroup       `json:"studyGroup"`
	Poster      *User             `json:"poster"`
	Name        string            `json:"name"`
	Content     string            `json:"content"`
	Attachments []*PostAttachment `json:"attachments"`
}

func (Post) IsNode()            {}
func (this Post) GetID() string { return this.ID }

type PostAttachment struct {
	ID          string  `json:"id"`
	UploadUUID  string  `json:"uploadUUID"`
	Description *string `json:"description,omitempty"`
	FileName    string  `json:"fileName"`
}

func (PostAttachment) IsNode()            {}
func (this PostAttachment) GetID() string { return this.ID }
