package main

import (
  "fmt"

  id "github.com/multiverse-os/muid"
)

func main() {
  fmt.Println("determinsitic id example")
  fmt.Println("========================")
  fmt.Println("  generate 3 ids from the same source string and verify they are the same")
  fmt.Println("  convert each back to the id using the resulting string, short (string), and noprefix (string)")

  deterministicId, err := id.FromString("test")
  if err != nil {
    panic(err)
  }
  fmt.Println(deterministicId.String())

  deterministicId, err = id.FromString("test")
  if err != nil {
    panic(err)
  }
  fmt.Println(deterministicId.String())

  deterministicId, err = id.FromString("test")
  if err != nil {
    panic(err)
  }
  fmt.Println(deterministicId.String())
}
