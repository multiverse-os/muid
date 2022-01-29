package muid

import (
  "crypto/sha3"
)

type HashType uint8

const (
  Undefined HashType = iota
  Shake256 
)

func HashBytes(hashType HashType, bytes []byte, length int) []byte {
  switch hashType {
  case Shake256:
  	shake := sha3.NewShake256()
		shake.Write(bytes)
    hash := make([]byte, length)
    _, err := shake.Read(hash)
    return hash
  default: // Undefined
    return bytes
  }

}
