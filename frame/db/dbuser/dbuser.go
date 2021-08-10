package dbuser

import (
	"database/sql"

	"github.com/covrom/hexagonarch/core/relation"
	"github.com/covrom/hexagonarch/core/user"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (us *Users) UserByID(uid relation.UserID) (*user.User, error) {
	// select from database
	return &user.User{}, nil
}

func (us *Users) SaveUser(u *user.User) error {
	// replace in database
	return nil
}
