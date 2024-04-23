// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type Node interface {
	IsNode()
	GetID() string
}

type AuthInfo struct {
	Provider AuthProviders `json:"provider"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
}

type ClassInfo struct {
	UniversityName string `json:"university_name"`
	Name           string `json:"name"`
	Description    string `json:"description"`
}

type Mutation struct {
}

type Profile struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	School         string `json:"school"`
	GraduationYear int    `json:"graduationYear"`
}

type Query struct {
}

type UploadWithMeta struct {
	Filename    string         `json:"filename"`
	Description *string        `json:"description,omitempty"`
	Upload      graphql.Upload `json:"upload"`
}

type AuthProviders string

const (
	AuthProvidersGithub           AuthProviders = "GITHUB"
	AuthProvidersGoogle           AuthProviders = "GOOGLE"
	AuthProvidersUsernamepassword AuthProviders = "USERNAMEPASSWORD"
)

var AllAuthProviders = []AuthProviders{
	AuthProvidersGithub,
	AuthProvidersGoogle,
	AuthProvidersUsernamepassword,
}

func (e AuthProviders) IsValid() bool {
	switch e {
	case AuthProvidersGithub, AuthProvidersGoogle, AuthProvidersUsernamepassword:
		return true
	}
	return false
}

func (e AuthProviders) String() string {
	return string(e)
}

func (e *AuthProviders) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuthProviders(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuthProviders", str)
	}
	return nil
}

func (e AuthProviders) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
