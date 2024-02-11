package helpers

import "github.com/google/uuid"

func GenerateUUID() string {
	randomUUID := uuid.New()
	return randomUUID.String()
}
