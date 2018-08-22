package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Encryption(t *testing.T) {
	c, err := newCrypto([]byte(strings.Repeat("a", 16)))
	assert.NoError(t, err)
	text := []byte("I am a plain text.")

	eOut, err := c.encrypt(text)
	assert.NoError(t, err)

	dOut, err := c.decrypt(eOut)
	assert.NoError(t, err)

	assert.Equal(t, text, dOut)
}
