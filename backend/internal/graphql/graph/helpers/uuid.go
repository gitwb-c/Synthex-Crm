package helpers

import "github.com/google/uuid"

func UuidParser(ids []string) ([]uuid.UUID, error) {
	var uuids []uuid.UUID

	for _, id := range ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			return []uuid.UUID{}, err
		}
		uuids = append(uuids, uid)
	}
	return uuids, nil
}
