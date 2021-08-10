package dbspace

import (
	"database/sql"

	"github.com/covrom/hexagonarch/core/relation"
	"github.com/covrom/hexagonarch/core/space"
)

type Spaces struct {
	db *sql.DB
}

func NewSpaces(db *sql.DB) *Spaces {
	return &Spaces{
		db: db,
	}
}

func (us *Spaces) SpaceByID(spid relation.SpaceID) (*space.Space, error) {
	// select from database
	return &space.Space{}, nil
}

func (us *Spaces) SaveSpace(sp *space.Space) error {
	// replace in database
	return nil
}

func (us *Spaces) MigrateSpace(sp *space.Space) error {
	// create database objects for space,i.e. CREATE SCHEMA
	return nil
}
