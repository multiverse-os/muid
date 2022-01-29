package muid

import (
  "encoding/binary"
  "math/rand"
)

func randomByte() byte {
  byteBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(byteBuffer, rand.Uint32())
  return byteBuffer[0]
}

func randomBytes(length int) (randomBytes []byte) {
  for index := 0; index < length; index++ {
    randomBytes = append(randomBytes, randomByte())
  }
  return randomBytes
}
