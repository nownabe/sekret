package main

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
)

type editCommand struct {
	*editorCommand
}

func editCommandFromContext(c *cli.Context) (*editCommand, error) {
	ecmd, err := editorCommandFromContext(c)
	if err != nil {
		return nil, err
	}

	if !exists(ecmd.filename) {
		return nil, fmt.Errorf("%s does not exist", ecmd.filename)
	}

	return &editCommand{ecmd}, nil
}

func (c *editCommand) run() error {
	cipherText, err := ioutil.ReadFile(c.filename)
	if err != nil {
		return err
	}

	plainText, err := c.crypto.decrypt(cipherText)
	if err != nil {
		return err
	}

	if c.decode {
		decoded, err := decode(plainText)
		if err != nil {
			return err
		}
		plainText = decoded
	}

	updatedPlainText, err := c.editText(plainText)
	if err != nil {
		return err
	}

	if c.decode {
		encoded, err := encode(updatedPlainText)
		if err != nil {
			return err
		}
		updatedPlainText = encoded
	}

	if err := c.validator.validate(updatedPlainText); err != nil {
		return err
	}

	return c.update(updatedPlainText)
}
