package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDncryptCommand_run(t *testing.T) {
	cmd := newTestCommand()
	dc := &decryptCommand{cmd}

	plainText := []byte("plaintext")

	cipherText, err := cmd.crypto.encrypt(plainText)
	if err != nil {
		panic(err)
	}

	in := bytes.NewBuffer(cipherText)
	out := &bytes.Buffer{}

	if assert.NoError(t, dc.run(in, out)) {
		assert.Equal(t, plainText, out.Bytes())
	}
}
