package main

import (
  "fmt"
  "os"
  "log"

  id "github.com/multiverse-os/muid"
	cli "github.com/multiverse-os/cli"
)

func main() {
  cmd := cli.New(&cli.CLI{
    Name: "muid",
    Version: cli.Version{Major: 0, Minor: 1, Patch: 0},
    Description: "A command-line tool for generating muid keys.",
    GlobalFlags: []cli.Flag{
      cli.Flag{
        Name:        "randbytes",
        Alias:       "r",
        Default:     "3",
        Description: "Specify number of random bytes to use",
      },
    },
  })

  _, err := cmd.Parse(os.Args)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Print(id.Deterministic("seed", 12).Base32().String(), "\n")
  fmt.Print(id.Deterministic("seed", 12).Base32().String(), "\n")
  // TODO: Add flags for various key options using multiverse-os/cmd framework
}
