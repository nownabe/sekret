package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

func encryptCommand(c *cli.Context) error {
	key, err := keyFromContext(c)
	if err != nil {
		return err
	}

	if c.NArg() != 1 {
		return fmt.Errorf("file is required")
	}

	filename := c.Args()[0]
	if !exists(filename) {
		return fmt.Errorf("%s does not exist", filename)
	}

	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	return encrypt(key, in, os.Stdout)
}

func encrypt(key []byte, in io.Reader, out io.Writer) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plaintext, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	ciphertext = append(nonce, ciphertext...)
	_, err = out.Write(ciphertext)
	if err != nil {
		return err
	}

	return nil
}
