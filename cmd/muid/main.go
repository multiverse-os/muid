package main

import (
  "fmt"

  id "github.com/multiverse-os/muid"
)

func main() {
  fmt.Println("[muid generator] not yet implemented")
  fmt.Println("output should just be the result so it can be used inline easily in scripting langauges")
  fmt.Println("    [", fmt.Sprintf("%v", id.Generate().Base32().String()), "]")
}
