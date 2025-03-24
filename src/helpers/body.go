package helpers

import (
	"encoding/json"
	"io"
)

func ParseBody[T any](body io.ReadCloser) (*T, error) {
	var t T
	err := json.NewDecoder(body).Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
