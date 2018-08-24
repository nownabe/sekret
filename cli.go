package main

import (
	"os"

	"github.com/urfave/cli"
)

const (
	cliName = "sekret"
	cliDesc = "Encrypt/Decrypt Kubernetes Secrets"
	cliVer  = "1.0.0"

	decodeBase64Flagname = "decode-base64"
	editorFlagName       = "editor"
	keyFlagName          = "key"
)

func newCLI() *cli.App {
	app := cli.NewApp()

	app.Name = cliName
	app.Usage = cliDesc
	app.Version = cliVer

	keyFlag := cli.StringFlag{
		Name:   keyFlagName,
		Usage:  "Encryption key",
		EnvVar: "ENCRYPTION_KEY",
	}

	app.Flags = []cli.Flag{keyFlag}

	app.Commands = []cli.Command{
		{
			Name:      "edit",
			Usage:     "Edit secrets as plain text",
			ArgsUsage: "file",
			Flags: []cli.Flag{
				keyFlag,
				cli.BoolTFlag{
					Name:  decodeBase64Flagname,
					Usage: "Decode base64 data",
				},
				cli.StringFlag{
					Name:   editorFlagName,
					Usage:  "Editor",
					EnvVar: "EDITOR",
				},
			},
			Action: editAction,
		},
		{
			Name:      "encrypt",
			ShortName: "enc",
			Usage:     "Encrypt file",
			ArgsUsage: "file",
			Flags:     []cli.Flag{keyFlag},
			Action:    encryptAction,
		},
		{
			Name:      "decrypt",
			ShortName: "dec",
			Usage:     "Decrypt encrypted file",
			ArgsUsage: "file",
			Flags:     []cli.Flag{keyFlag},
			Action:    decryptAction,
		},
	}

	return app
}

func editAction(c *cli.Context) error {
	cmd, err := editCommandFromContext(c)
	if err != nil {
		return err
	}
	return cmd.run()
}

func encryptAction(c *cli.Context) error {
	cmd, err := encryptCommandFromContext(c)
	if err != nil {
		return err
	}

	in, err := os.Open(cmd.filename)
	if err != nil {
		return err
	}
	defer in.Close()

	return cmd.run(in, os.Stdout)
}

func decryptAction(c *cli.Context) error {
	cmd, err := decryptCommandFromContext(c)
	if err != nil {
		return err
	}

	in, err := os.Open(cmd.filename)
	if err != nil {
		return err
	}
	defer in.Close()

	return cmd.run(in, os.Stdout)
}
