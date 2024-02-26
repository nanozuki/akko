package service

import (
	"context"
	"errors"
	"net/http"
)

type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var users = []*User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Lily Bush"},
}

var tokens = map[string]int{
	"affe1e5c-be35-43c3-9ef5-46cd1f4a0d2c": 1,
	"4b0ca88d-ffb5-49b4-9d7d-c666e2f2d3c2": 2,
}

func LoadUserByToken(req *http.Request) (*User, error) {
	token := req.Header.Get("Token")
	if userID, ok := tokens[token]; ok {
		for _, user := range users {
			if user.ID == userID {
				return user, nil
			}
		}
	}
	return nil, errors.New("unauthorized")
}

// GetUser returns logged in user.
// @get /users
func (s *Service) GetUser(ctx context.Context, user *User) (*User, error) {
	return user, nil
}

type UserPatch struct {
	Name *string `json:"name,omitempty"`
}

// PatchUser updates logged in user.
// @patch /users/{id}
func (s *Service) PatchUser(ctx context.Context, user *User, patch UserPatch) (*User, error) {
	if patch.Name != nil {
		user.Name = *patch.Name
	}
	return user, nil
}
