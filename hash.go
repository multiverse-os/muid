package muid

import (
  sha3 "golang.org/x/crypto/sha3"
)

type HashType uint8

const (
  NoHash HashType = iota
  SHA3_224
  Shake128
  Shake256
)

// TODO: The (Id) Hash() Id chainable to hash any key 

func hashBytes(hashType HashType, bytes []byte, length int) []byte {
  idHash := make([]byte, length)
  switch hashType {
  case SHA3_224:
    hash := sha3.Sum224(bytes)
    return hash[:]
  case Shake128:
    sha3.ShakeSum256(idHash, bytes)
    return idHash
  default: // Undefined
    return bytes
  }

}
