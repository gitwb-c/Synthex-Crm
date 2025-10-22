package utils

import (
	"fmt"

	"github.com/google/uuid"
)

// ParseUUIDs pode aceitar tanto um único ID como uma lista de IDs.
// Se um único ID for fornecido, ele retornará um ponteiro para UUID. Se uma lista for fornecida, ele retornará uma lista de ponteiros para UUIDs.
func ParseUUIDs(ids interface{}) (interface{}, error) {
	switch v := ids.(type) {
	case string:
		// Se for um único ID, convertemos diretamente para *uuid.UUID
		uid, err := uuid.Parse(v)
		if err != nil {
			return nil, fmt.Errorf("invalid UUID format for ID %s: %v", v, err)
		}
		return &uid, nil // Retorna o ponteiro para o UUID

	case []string:
		var uuidList []*uuid.UUID
		// Se for uma lista de IDs, convertemos cada um e retornamos uma lista de ponteiros
		for _, id := range v {
			uid, err := uuid.Parse(id)
			if err != nil {
				return nil, fmt.Errorf("invalid UUID format for ID %s: %v", id, err)
			}
			uuidList = append(uuidList, &uid) // Adiciona o ponteiro para a lista
		}
		return uuidList, nil

	default:
		return nil, fmt.Errorf("expected string or []string, got %T", v)
	}
}
