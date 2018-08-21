package main

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Encryption(t *testing.T) {
	key := []byte(strings.Repeat("a", 16))
	text := []byte("I am a plain text.")

	eIn := bytes.NewBuffer(text)
	eOut := new(bytes.Buffer)
	dOut := new(bytes.Buffer)

	if assert.NoError(t, encrypt(key, eIn, eOut)) {
		if assert.NoError(t, decrypt(key, eOut, dOut)) {
			decText, _ := ioutil.ReadAll(dOut)
			assert.Equal(t, text, decText)
		}
	}
}
