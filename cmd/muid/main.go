package main

import (
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli"
	id "github.com/multiverse-os/muid"
)

func main() {
	cmd, initErrors := cli.New(cli.App{
		Name:        "muid",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 0},
		Description: "A command-line tool for generating muid keys.",
		GlobalFlags: cli.Flags(
			cli.Flag{
				Name:        "random",
				Alias:       "r",
				Default:     "3",
				Description: "Specify number of random bytes to use",
			},
			cli.Flag{
				Name:        "append",
				Alias:       "a",
				Default:     "",
				Description: "Append text as prefix or suffix to the `muid`",
			},
			cli.Flag{
				Name:        "time",
				Alias:       "t",
				Default:     "now",
				Description: "Time used to seed the `muid`",
			},
			cli.Flag{
				Name:        "checksum",
				Alias:       "c",
				Default:     "simple",
				Description: "Hash `muid` using specified checksum alogrithm",
			},
			cli.Flag{
				Name:        "compress",
				Alias:       "x",
				Default:     "", // TODO: Should be bool
				Description: "Compress `muid` using specified algorithm",
			},
			cli.Flag{
				Name:        "encoding",
				Alias:       "e",
				Default:     "", // Base32+Hex, Base32, Hex, Bytes, String, Base64, Base58, etc
				Description: "Encode `muid` using specified algorithm",
			},
			cli.Flag{
				Name:        "machine",
				Alias:       "m",
				Default:     "3", // TODO: Should be uint8
				Description: "Machine Id bytes to be included in `muid`",
			},
			cli.Flag{
				Name:        "pid",
				Alias:       "p",
				Default:     "2", // TODO: Should be uint8
				Description: "Prociess Id (PID) bytes to be included in `muid`",
			},
			cli.Flag{
				Name:        "compress",
				Alias:       "x",
				Default:     "",
				Description: "Compression algorithm used on the `muid`",
			},
			cli.Flag{
				Name:        "help",
				Alias:       "h",
				Default:     "", // TODO: Should be bool
				Description: "Commands, sub-commands, flags, and descriptions of `muid` and featuers",
			},
			cli.Flag{
				Name:        "version",
				Alias:       "v",
				Default:     "", // TODO: Should be bool
				Description: "Show the current build version of `muid` command-line tool",
			},
		),
	})

	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	}

	fmt.Print(id.Generate().Base32().String(), "\n")
	//fmt.Print(id.Deterministic("seed-address-02-02-20", 6).Base32().String(), "\n")
}
