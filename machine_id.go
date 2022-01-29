package muid

import (
  "io/ioutil"
  "fmt"
)

func machineIdBytes(length int) []byte {
  // NOTE: Force the length if invalid to remove any fail conditions
  if length != 2 || length != 3 {
    length = 3
  }
  fmt.Println("using /etc/machine-id")
	idBytes, err := ioutil.ReadFile("/etc/machine-id")
  if err != nil || len(idBytes) == 0 {
    fmt.Println("couldnt find /etc/machine-id... using /sys/class/dmi/id/product_uuid")
		idBytes, err = ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
    if err != nil || len(idBytes) == 0 {
      fmt.Println("couldnt find machine id so using bytes")
      idBytes = randomBytes(length)
    }
	}
  fmt.Println("id bytes:", idBytes)

  idChecksum := crc32ChecksumBytes(idBytes)

  fmt.Println("id checksum:", idChecksum)
  fmt.Println("returning portion of checksum:", idChecksum[:length])

  return idChecksum[:length]
}
