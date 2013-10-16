package main

import (
	"crypto/rand"
	"io"
)

func makeKey(strength int) ([]byte, error) {
	b := make([]byte, strength)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
