package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptCommand_run(t *testing.T) {
	cmd := newTestCommand()
	ec := &encryptCommand{cmd}

	in := bytes.NewBufferString("plaintext")
	out := &bytes.Buffer{}

	if assert.NoError(t, ec.run(in, out)) {
		assert.NotEmpty(t, out.Bytes())
	}
}
