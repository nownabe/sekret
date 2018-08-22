package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli"
)

func editCommand(c *cli.Context) error {
	key, err := keyFromContext(c)
	if err != nil {
		return err
	}

	editor := c.String(editorFlagName)
	if editor == "" {
		return fmt.Errorf("editor is required")
	}

	if c.NArg() != 1 {
		return fmt.Errorf("file is required")
	}

	filename := c.Args()[0]
	if !exists(filename) {
		return fmt.Errorf("%s does not exist", filename)
	}

	cipherText, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	cipherTextBuf := bytes.NewBuffer(cipherText)

	plainTextBuf := new(bytes.Buffer)
	if err := decrypt(key, cipherTextBuf, plainTextBuf); err != nil {
		return err
	}

	tmpfile, err := ioutil.TempFile("", path.Base(filename))
	if err != nil {
		return err
	}
	defer removeTempFile(tmpfile.Name())

	if _, err := tmpfile.Write(plainTextBuf.Bytes()); err != nil {
		return err
	}

	cmd := exec.Command(editor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	updatedPlainText, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		return err
	}
	updatedPlainTextBuf := bytes.NewBuffer(updatedPlainText)

	updatedCipherTextBuf := new(bytes.Buffer)
	if err := encrypt(key, updatedPlainTextBuf, updatedCipherTextBuf); err != nil {
		return err
	}

	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}

	tmpEncFile := filename + ".tmp"
	if err := ioutil.WriteFile(tmpEncFile, updatedCipherTextBuf.Bytes(), fi.Mode()); err != nil {
		return err
	}
	defer removeTempFile(tmpEncFile)

	if err := os.Rename(tmpEncFile, filename); err != nil {
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
