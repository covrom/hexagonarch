package relation

type UserIDs []UserID

func (ids *UserIDs) Add(uid UserID) {
	fnd := false
	for _, v := range *ids {
		if v == uid {
			fnd = true
			break
		}
	}
	if !fnd {
		*ids = append(*ids, uid)
	}
}
