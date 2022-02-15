package muid

import (
  sha3 "golang.org/x/crypto/sha3"
)

// TODO: We are looking for a two-way hash was can easily use to convert forward
// and back, preferably compressed afterwards to get a small byte length for our
// final byte slice deterministic id. However we just need need to get one
// working and build it out using interfaces so we can easily extend the hashing
// algorithm to any relevant one or any community desired one. The key here is,
// focus on the broad, and come back and decide the specific later by making it
// simple with a solid broad hashing system. 

// TODO: Defaulting to SHA3 because of the length it gives us and the general
//       respect for the standard. But this is explicitly designed to easily
//       support more; like xxHash which I may add soon. But this is explicitly
//       intended to to be easily modified by the developer to fit their
//       specific purposes with little effort, and each portion of the id system
//       works in this modular way.

type HashType uint8

const (
  NoHash HashType = iota
  SHA3
  Shake128
  Shake256
)


// TODO: The (Id) Hash() Id chainable to hash any key 
func hashBytes(hashType HashType, bytes []byte) []byte {
  var idHash []byte
  switch hashType {
  case SHA3:
    hash := sha3.Sum224(bytes)
    return hash[:]
  case Shake128:
    sha3.ShakeSum256(idHash, bytes)
    return idHash
  default: // Undefined
    return bytes
  }
}
