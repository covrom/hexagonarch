package space

import (
	"fmt"

	"github.com/covrom/hexagonarch/core/relation"
)

type SpaceStorer interface {
	SpaceByID(spid relation.SpaceID) (*Space, error)
	SaveSpace(sp *Space) error
	MigrateSpace(sp *Space) error
}

type Spaces struct {
	spaceStore SpaceStorer
}

func NewSpaces(st SpaceStorer) *Spaces {
	return &Spaces{
		spaceStore: st,
	}
}

func (us *Spaces) SaveSpace(sp *Space) error {
	isNew := sp.Id.IsEmpty()
	if isNew {
		sp.Id = relation.NewSpaceID()
	}
	if sp.Name == "" {
		return fmt.Errorf("register space error: name is empty")
	}
	if err := us.spaceStore.SaveSpace(sp); err != nil {
		return fmt.Errorf("add space %q error: %w", sp.Name, err)
	}
	if isNew {
		if err := us.spaceStore.MigrateSpace(sp); err != nil {
			return fmt.Errorf("add space %q error: %w", sp.Name, err)
		}
	}
	return nil
}

func (us *Spaces) SpaceByID(spid relation.SpaceID) (*Space, error) {
	if spid.IsEmpty() {
		return nil, fmt.Errorf("space id is empty")
	}
	sp, err := us.spaceStore.SpaceByID(spid)
	if err != nil {
		return nil, fmt.Errorf("get space by id %q error: %w", spid, err)
	}
	return sp, nil
}
