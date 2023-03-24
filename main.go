package main

import (
	"os"

	"github.com/vinhph0906/blockchain-wallet/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()

}
