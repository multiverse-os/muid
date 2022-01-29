package muid

import (
  "crypto/rand"
)

func RandomByte() byte {
	return byte(rand.Uint8())
}

func RandomBytes(length uint8) (randomBytes []byte) {
  for index := 0; index < length; index++ {
    randomBytes = append(randomBytes, RandomByte())
  }
  return randomBytes
}

func RandomUint32() uint32 {
  return rand.Uint32()
}
