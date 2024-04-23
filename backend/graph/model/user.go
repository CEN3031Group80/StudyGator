package model

type FriendRequest struct {
	ID       string `json:"id"`
	Sender   *User  `json:"sender"`
	Receiver *User  `json:"receiver"`
	Accepted bool   `json:"accepted"`
}

func (FriendRequest) IsNode()            {}
func (this FriendRequest) GetID() string { return this.ID }

type User struct {
	ID        string    `json:"id"`
	AvatarURL string    `json:"avatarURL"`
	AuthInfo  *AuthInfo `json:"authInfo"`
	Profile   *Profile  `json:"profile,omitempty"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }
