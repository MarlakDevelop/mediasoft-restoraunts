package uuid

import "github.com/google/uuid"

func FromStrings(strings []string) []uuid.UUID {
	uuids := make([]uuid.UUID, 0, len(strings))
	for _, s := range strings {
		uuid_, err := uuid.Parse(s)
		if err != nil {
			continue
		}
		uuids = append(uuids, uuid_)
	}
	return uuids
}

func ToStrings(uuids []uuid.UUID) []string {
	strings := make([]string, 0, len(uuids))
	for _, uuid_ := range uuids {
		strings = append(strings, uuid_.String())
	}
	return strings
}
