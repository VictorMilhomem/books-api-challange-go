package models

import "github.com/google/uuid"

type AuthorName struct {
	Name string
}

type Author struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewAuthorName(name string) *AuthorName {
	return &AuthorName{
		Name: name,
	}
}
