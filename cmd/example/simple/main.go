package main

import (
	"fmt"

	id "github.com/multiverse-os/muid"
)

func main() {
	fmt.Println("multiverse sortable time stamped id")
	fmt.Println("============================================================")

  id1 := id.Generate()
  id2 := id.Generate()

  fmt.Println("id 1 is: ", id1)
  fmt.Println("id 2 is: ", id2)

  fmt.Println("now testing comparison")

  fmt.Println(" id1.Compare(id1) = ", id1.Compare(id1))
  fmt.Println(" id1.Compare(id2) = ", id1.Compare(id2))
  fmt.Println(" id1.Compare(id1) = ", id2.Compare(id1))
  fmt.Println(" id2.Compare(id2) = ", id2.Compare(id2))

  fmt.Println("now converting to base32 and comparing")

  id1 = id1.Base32()
  id2 = id2.Base32()

  fmt.Println("id (base32) 1 is: ", id1)
  fmt.Println("id (base32) 2 is: ", id2)

  fmt.Println(" id1.Compare(id1) = ", id1.Compare(id1))
  fmt.Println(" id1.Compare(id2) = ", id1.Compare(id2))
  fmt.Println(" id2.Compare(id1) = ", id2.Compare(id1))
  fmt.Println(" id2.Compare(id2) = ", id2.Compare(id2))
}
