package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli"
)

type editCommand struct {
	*command
	decode    bool
	editor    string
	validator *validator
}

func editCommandFromContext(c *cli.Context) (*editCommand, error) {
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

	return &editCommand{
		cmd,
		c.Bool(decodeBase64Flagname),
		editor,
		validator,
	}, nil
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

func (c *editCommand) update(plainText []byte) error {
	tmpFile, err := ioutil.TempFile("", path.Base(c.filename))
	if err != nil {
		return err
	}
	defer removeTempFile(tmpFile.Name())

	cipherText, err := c.crypto.encrypt(plainText)
	if err != nil {
		return err
	}

	fi, err := os.Stat(c.filename)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(tmpFile.Name(), cipherText, fi.Mode()); err != nil {
		return err
	}

	if err := os.Rename(tmpFile.Name(), c.filename); err != nil {
		return err
	}

	return nil
}

func (c *editCommand) editText(text []byte) ([]byte, error) {
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

func removeTempFile(path string) {
	if exists(path) {
		if err := os.Remove(path); err != nil {
			panic(err)
		}
	}
}
