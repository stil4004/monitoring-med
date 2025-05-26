package utils

import "github.com/google/uuid"

func GenerateHash() string {
	newUUID := uuid.New()
	return newUUID.String()
}
