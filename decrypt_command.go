package main

import (
	"fmt"
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

	if !exists(cmd.filename) {
		return nil, fmt.Errorf("%s does not exist", cmd.filename)
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

	if _, err := out.Write(plainText); err != nil {
		return err
	}

	return nil
}
