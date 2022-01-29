package muid

import (
  "encoding/binary"
  "math/rand"
)

func RandomByte() byte {
  byteBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(byteBuffer, rand.Uint32())
  return byteBuffer[0]
}

func RandomBytes(length uint8) (randomBytes []byte) {
  for index := 0; index < int(length); index++ {
    randomBytes = append(randomBytes, RandomByte())
  }
  return randomBytes
}
