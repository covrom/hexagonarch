package userspace

import (
	"fmt"

	"github.com/covrom/hexagonarch/core/relation"
	"github.com/covrom/hexagonarch/core/space"
	"github.com/covrom/hexagonarch/core/user"
)

type UserSpace struct {
	users  *user.Users
	spaces *space.Spaces
}

func NewUserSpace(users *user.Users, spaces *space.Spaces) *UserSpace {
	return &UserSpace{
		users:  users,
		spaces: spaces,
	}
}

func (us *UserSpace) AssignUserToSpace(uid relation.UserID, spid relation.SpaceID) error {
	u, err := us.users.UserByID(uid)
	if err != nil {
		return fmt.Errorf("assign user to space error when searching for user by ID %q : %w", uid, err)
	}
	sp, err := us.spaces.SpaceByID(spid)
	if err != nil {
		return fmt.Errorf("assign user to space error when searching for space by ID %q : %w", spid, err)
	}

	u.SpaceIDs.Add(sp.Id)
	sp.UserIDs.Add(u.Id)

	if err := us.users.SaveUser(u); err != nil {
		return fmt.Errorf("save user error : %w", err)
	}
	if err := us.spaces.SaveSpace(sp); err != nil {
		return fmt.Errorf("save space error : %w", err)
	}
	return nil
}

func (us *UserSpace) RegisterUserInSpaces(u *User) error {
	cu := &user.User{
		Name:     u.Name,
		Login:    u.Login,
		Password: u.Password,
		Email:    u.Email,
	}
	if err := us.users.SaveUser(cu); err != nil {
		return fmt.Errorf("register user in spaces saving user error: %w", err)
	}
	for _, spid := range u.SpaceIDs {
		if err := us.AssignUserToSpace(cu.Id, spid); err != nil {
			return fmt.Errorf("assign user to space error when register user in spaces: %w", err)
		}
	}
	return nil
}

func (us *UserSpace) CreateSpace(sp Space) (*Space, error) {
	newsp := &space.Space{
		Name: sp.Name,
	}
	if err := us.spaces.SaveSpace(newsp); err != nil {
		return nil, fmt.Errorf("create space error: %w", err)
	}

	return &Space{
		Name: sp.Name,
		ID:   newsp.Id,
	}, nil
}
