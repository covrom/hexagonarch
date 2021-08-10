package relation

import "github.com/google/uuid"

type UserID uuid.UUID

func NewUserID() UserID {
	return UserID(uuid.New())
}

func (uid UserID) IsEmpty() bool {
	return uuid.UUID(uid) == uuid.Nil
}
