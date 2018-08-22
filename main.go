package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const (
	cliName = "sekret"
	cliDesc = "Encrypt/Decrypt Kubernetes Secrets"
	cliVer  = "0.1.0"

	keyFlagName = "key"
)

func newApp() *cli.App {
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
					Name:  "decode-base64",
					Usage: "Decode base64 data",
				},
			},
			Action: editCommand,
		},
		{
			Name:      "encrypt",
			ShortName: "enc",
			Usage:     "Encrypt file",
			ArgsUsage: "file",
			Flags:     []cli.Flag{keyFlag},
			Action:    encryptCommand,
		},
		{
			Name:      "decrypt",
			ShortName: "dec",
			Usage:     "Decrypt encrypted file",
			ArgsUsage: "file",
			Flags:     []cli.Flag{keyFlag},
			Action:    decryptCommand,
		},
	}

	return app
}

func main() {
	app := newApp()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
