package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type crypto struct {
	gcm cipher.AEAD
}

func newCrypto(key []byte) (*crypto, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &crypto{gcm: gcm}, nil
}

func (c *crypto) encrypt(in []byte) ([]byte, error) {
	nonce := make([]byte, c.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err

	}

	cipherText := c.gcm.Seal(nil, nonce, in, nil)
	cipherText = append(nonce, cipherText...)
	return cipherText, nil
}

func (c *crypto) decrypt(in []byte) ([]byte, error) {
	nonce := in[:c.gcm.NonceSize()]
	return c.gcm.Open(nil, nonce, in[c.gcm.NonceSize():], nil)
}
