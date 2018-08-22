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
	editor string
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

	return &editCommand{cmd, editor}, nil
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

	tmpfile, err := ioutil.TempFile("", path.Base(c.filename))
	if err != nil {
		return err
	}
	defer removeTempFile(tmpfile.Name())

	if _, err := tmpfile.Write(plainText); err != nil {
		return err
	}

	cmd := exec.Command(c.editor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	updatedPlainText, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		return err
	}

	updatedCipherText, err := c.crypto.encrypt(updatedPlainText)
	if err != nil {
		return err
	}

	fi, err := os.Stat(c.filename)
	if err != nil {
		return err
	}

	tmpEncFile := c.filename + ".tmp"
	if err := ioutil.WriteFile(tmpEncFile, updatedCipherText, fi.Mode()); err != nil {
		return err
	}
	defer removeTempFile(tmpEncFile)

	if err := os.Rename(tmpEncFile, c.filename); err != nil {
		return err
	}

	return nil

}

func removeTempFile(path string) {
	if exists(path) {
		if err := os.Remove(path); err != nil {
			panic(err)
		}
	}
}
