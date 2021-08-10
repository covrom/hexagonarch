package userspace

import "github.com/covrom/hexagonarch/core/relation"

type Space struct {
	Name string           `json:"name"`
	ID   relation.SpaceID `json:"id"`
}
