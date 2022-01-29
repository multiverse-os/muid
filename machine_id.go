package muid

import (
  "io/ioutil"
)

func machineIdBytes(length int) []byte {
	idBytes, err := ioutil.ReadFile("/etc/machine-id")
  if err != nil || len(idBytes) == 0 {
		idBytes, err = ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
    if err != nil || len(idBytes) == 0 {
      idBytes = randomBytes(length)
    }
	}
  idChecksum := crc32ChecksumBytes(idBytes)

  return idChecksum[:length]
}
