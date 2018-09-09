package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/urfave/cli"
)

type encryptCommand struct {
	*command
}

func encryptCommandFromContext(c *cli.Context) (*encryptCommand, error) {
	cmd, err := commandFromContext(c)
	if err != nil {
		return nil, err
	}

	if !exists(cmd.filename) {
		return nil, fmt.Errorf("%s does not exist", cmd.filename)
	}

	return &encryptCommand{cmd}, nil
}

func (c *encryptCommand) run(in io.Reader, out io.Writer) error {
	plainText, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	cipherText, err := c.crypto.encrypt(plainText)
	if err != nil {
		return err
	}

	if _, err := out.Write(cipherText); err != nil {
		return err
	}

	return nil
}
