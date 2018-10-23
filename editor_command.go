package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli"
)

type editorCommand struct {
	*command
	decode    bool
	editor    string
	validator *validator
}

func editorCommandFromContext(c *cli.Context) (*editorCommand, error) {
	cmd, err := commandFromContext(c)
	if err != nil {
		return nil, err
	}

	editor := c.String(editorFlagName)
	if editor == "" {
		return nil, fmt.Errorf("editor is required")
	}

	validator, err := newValidator()
	if err != nil {
		return nil, err
	}
	return &editorCommand{
		cmd,
		c.Bool(decodeBase64FlagName),
		editor,
		validator,
	}, nil
}

func (c *editorCommand) create(plainText []byte) error {
	cipherText, err := c.crypto.encrypt(plainText)
	if err != nil {
		return err
	}

	// re-check before save
	if exists(c.filename) {
		return fmt.Errorf("%s already exists", c.filename)
	}

	return ioutil.WriteFile(c.filename, cipherText, 0644)
}

func (c *editorCommand) update(plainText []byte) error {
	cipherText, err := c.crypto.encrypt(plainText)
	if err != nil {
		return err
	}

	fi, err := os.Stat(c.filename)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.filename, cipherText, fi.Mode())
}

func (c *editorCommand) editText(text []byte) ([]byte, error) {
	tmpFile, err := ioutil.TempFile("", path.Base(c.filename))
	if err != nil {
		return nil, err
	}
	defer removeTempFile(tmpFile.Name())

	if _, err := tmpFile.Write(text); err != nil {
		return nil, err
	}

	cmd := exec.Command(c.editor, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(tmpFile.Name())
}
