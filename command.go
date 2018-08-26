package main

import (
	"fmt"

	"github.com/urfave/cli"
)

type command struct {
	crypto   *crypto
	filename string
}

func commandFromContext(c *cli.Context) (*command, error) {
	key, err := keyFromContext(c)
	if err != nil {
		return nil, err
	}

	if c.NArg() != 1 {
		return nil, fmt.Errorf("file is required")
	}

	cr, err := newCrypto(key)
	if err != nil {
		return nil, err
	}

	return &command{
		crypto:   cr,
		filename: c.Args()[0],
	}, nil
}

func keyFromContext(c *cli.Context) ([]byte, error) {
	var key []byte
	if k := c.String(keyFlagName); k != "" {
		key = []byte(k)
	} else if k := c.GlobalString(keyFlagName); k != "" {
		key = []byte(k)
	} else {
		return nil, fmt.Errorf("key is required")
	}

	if len(key) != 16 && len(key) != 32 {
		return nil, fmt.Errorf("key must be 16 bytes or 32 bytes")
	}
	return key, nil
}
