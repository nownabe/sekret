package main

import (
	"fmt"
	"os"
)

func main() {
	cli := newCLI()
	err := cli.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
