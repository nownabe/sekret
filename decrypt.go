package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

func decryptCommand(c *cli.Context) error {
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

	return decrypt(key, in, os.Stdout)
}

func decrypt(key []byte, in io.Reader, out io.Writer) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	ciphertext, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	nonce := ciphertext[:gcm.NonceSize()]

	plaintext, err := gcm.Open(nil, nonce, ciphertext[gcm.NonceSize():], nil)
	if err != nil {
		return err
	}
	out.Write(plaintext)

	return nil
}
