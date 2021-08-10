package relation

type SpaceIDs []SpaceID

func (ids *SpaceIDs) Add(spid SpaceID) {
	fnd := false
	for _, v := range *ids {
		if v == spid {
			fnd = true
			break
		}
	}
	if !fnd {
		*ids = append(*ids, spid)
	}
}
