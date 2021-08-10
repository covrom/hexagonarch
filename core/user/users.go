package user

import (
	"fmt"

	"github.com/covrom/hexagonarch/core/relation"
)

type UserStorer interface {
	UserByID(uid relation.UserID) (*User, error)
	SaveUser(u *User) error
}

type Users struct {
	userStore UserStorer
}

func NewUsers(st UserStorer) *Users {
	return &Users{
		userStore: st,
	}
}

func (us *Users) SaveUser(u *User) error {
	if u.Id.IsEmpty() {
		u.Id = relation.NewUserID()
	}
	if u.Login == "" {
		return fmt.Errorf("save user error: login is empty")
	}
	if u.Email == "" {
		return fmt.Errorf("save user %q error: email is empty", u.Login)
	}
	if err := us.userStore.SaveUser(u); err != nil {
		return fmt.Errorf("save user %q error: %w", u.Login, err)
	}
	return nil
}

func (us *Users) UserByID(uid relation.UserID) (*User, error) {
	if uid.IsEmpty() {
		return nil, fmt.Errorf("user id is empty")
	}
	sp, err := us.userStore.UserByID(uid)
	if err != nil {
		return nil, fmt.Errorf("get user by id %q error: %w", uid, err)
	}
	return sp, nil
}
