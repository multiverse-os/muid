package muid

import (
  sha3 "golang.org/x/crypto/sha3"
)

type HashType uint8

const (
  Undefined HashType = iota
  Shake256 
)

// TODO: The (Id) Hash() Id chainable to hash any key 

func hashBytes(hashType HashType, bytes []byte, length int) []byte {
  switch hashType {
  case Shake256:
  	shake := sha3.NewShake256()
		shake.Write(bytes)
    hash := make([]byte, length)
    _, err := shake.Read(hash)
    if err != nil {
      return bytes
    }
    return hash
  default: // Undefined
    return bytes
  }

}
