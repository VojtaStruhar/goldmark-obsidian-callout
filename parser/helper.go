package parser

import (
	"errors"
)

func indexOf(byteArray []byte, target byte) (int, error) {
	for i, b := range byteArray {
		if b == target {
			return i, nil
		}
	}
	return -1, errors.New("index not found")
}
