package main

import (
	"context"

	"github.com/nanozuki/akko"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	return nil, nil
}

type User struct {
	gorm.Model
	Name    string
	Address string
}

// GetUserByID get user by id
// #[get(/user/<id>)]
func GetUserByID(ctx context.Context, db *gorm.DB, id int) (*User, error) {
	return nil, nil
}

// #[post(/user), data=<user>:json]
func CreateUser(ctx context.Context, db *gorm.DB, user *User) error {
	return nil
}

// #[get(/user?<name>)]
func ListUsers(ctx context.Context, db *gorm.DB, name []string) ([]*User, error) {
	return nil, nil
}

func AkkoBuild(a *akko.Akko) {
	a.Attach(akko.OnIgnite("connect db", func(a *akko.Akko) {
		a.Attach(akko.Provide(ConnectDB))
	}))
	a.Mount("/", []akko.Route{GetUserByID, CreateUser, ListUsers})
}
