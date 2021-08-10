package userspace

import "github.com/covrom/hexagonarch/core/relation"

type User struct {
	Name     string            `json:"name"`
	Login    string            `json:"login"`
	Password string            `json:"password"`
	Email    string            `json:"email"`
	SpaceIDs relation.SpaceIDs `json:"spaceIds"`
}
