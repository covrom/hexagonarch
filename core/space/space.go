package space

import "github.com/covrom/hexagonarch/core/relation"

type Space struct {
	Id      relation.SpaceID
	Name    string
	UserIDs relation.UserIDs
}
