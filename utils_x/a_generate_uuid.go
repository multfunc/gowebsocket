package utils_x

import (
	uuid "github.com/satori/go.uuid"
)

func GenerateUUID() (interface{}, error) {
	u := uuid.NewV1()
	return u.String(), nil
}
