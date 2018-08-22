package main

import (
	"io"
	"io/ioutil"

	"github.com/urfave/cli"
)

type decryptCommand struct {
	*command
}

func decryptCommandFromContext(c *cli.Context) (*decryptCommand, error) {
	cmd, err := commandFromContext(c)
	if err != nil {
		return nil, err
	}

	return &decryptCommand{cmd}, nil
}

func (c *decryptCommand) run(in io.Reader, out io.Writer) error {
	cipherText, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	plainText, err := c.crypto.decrypt(cipherText)
	if err != nil {
		return err
	}
	out.Write(plainText)

	return nil
}
