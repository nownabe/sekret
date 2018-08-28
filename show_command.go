package main

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
)

type showCommand struct {
	*command
	decode bool
}

func showCommandFromContext(c *cli.Context) (*showCommand, error) {
	cmd, err := commandFromContext(c)
	if err != nil {
		return nil, err
	}

	if !exists(cmd.filename) {
		return nil, fmt.Errorf("%s does not exist", cmd.filename)
	}

	return &showCommand{
		cmd,
		c.Bool(decodeBase64FlagName),
	}, nil
}

func (c *showCommand) run() error {
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

	fmt.Print(string(plainText))

	return nil
}
