package user

import "github.com/covrom/hexagonarch/core/relation"

type User struct {
	Id       relation.UserID
	Name     string
	Login    string
	Password string
	Email    string
	SpaceIDs relation.SpaceIDs
}
