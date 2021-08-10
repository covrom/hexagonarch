package relation

import "github.com/google/uuid"

type SpaceID uuid.UUID

func NewSpaceID() SpaceID {
	return SpaceID(uuid.New())
}

func (uid SpaceID) IsEmpty() bool {
	return uuid.UUID(uid) == uuid.Nil
}
