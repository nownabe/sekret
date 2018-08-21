package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

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

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
