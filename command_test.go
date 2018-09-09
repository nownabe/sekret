package main

import "strings"

var (
	testKey      = []byte(strings.Repeat("a", 16))
	testFilename = "testfile"
)

func newTestCommand() *command {
	cr, err := newCrypto(testKey)
	if err != nil {
		panic(err)
	}

	return &command{
		crypto:   cr,
		filename: testFilename,
	}
}
