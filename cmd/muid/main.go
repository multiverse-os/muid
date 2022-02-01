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
        Description: "Checksum type used in `muid`",
      },
      cli.Flag{
        Name:        "hash",
        Alias:       "h",
        Default:     "",
        Description: "Two-way hash to conceal information stored within `muid`",
      },
      cli.Flag{
        Name:        "compress",
        Alias:       "x",
        Default:     "",
        Description: "Compression algorithm used on the `muid`",
      },
    },
  })

  _, err := cmd.Parse(os.Args)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Print(id.Deterministic("seed", 12).Base32().String(), "\n")
  fmt.Print(id.Deterministic("seed", 12).Base32().Prefix("mv-").String(), "\n")
  fmt.Print(id.Deterministic("seed", 12).Base32().String(), "\n")
  // TODO: Add flags for various key options using multiverse-os/cmd framework
}
