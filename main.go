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
)

func main() {
	app := cli.NewApp()

	app.Name = cliName
	app.Usage = cliDesc
	app.Version = cliVer

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "key",
			Usage:  "Encryption key",
			EnvVar: "ENCRYPTION_KEY",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "edit",
			Usage:     "Edit secrets as plain text",
			ArgsUsage: "file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "key",
					Usage:  "Encryption key",
					EnvVar: "ENCRYPTION_KEY",
				},
				cli.BoolTFlag{
					Name:  "decode-base64",
					Usage: "Decode base64 data",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println(c.BoolT("decode-base64"))
				fmt.Println(c.String("key"))
				fmt.Println(c.Args())
				fmt.Println(c.GlobalString("key"))
				fmt.Println("edit")
				return nil
			},
		},
		{
			Name:      "encrypt",
			ShortName: "enc",
			Usage:     "Encrypt file",
			ArgsUsage: "file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "key",
					Usage:  "Encryption key",
					EnvVar: "ENCRYPTION_KEY",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println(c.Args())
				fmt.Println(c.String("key"))
				fmt.Println(c.GlobalString("key"))
				fmt.Println("encrypt")
				return nil
			},
		},
		{
			Name:      "decrypt",
			ShortName: "dec",
			Usage:     "Decrypt encrypted file",
			ArgsUsage: "file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "key",
					Usage:  "Encryption key",
					EnvVar: "ENCRYPTION_KEY",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println(c.Args())
				fmt.Println(c.String("key"))
				fmt.Println(c.GlobalString("key"))
				fmt.Println("decrypt")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
