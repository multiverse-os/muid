package muid

import (
  "encoding/hex"
)

func (self Id) Hex() string   { return hex.EncodeToString(self[:]) }
func (self Id) Bytes() []byte { return self[:] }

func (self Id) String() string {
	//text := make([]byte, stringEncodedLength)
	//encode(text, self[:])
	//return *(*string)(unsafe.Pointer(&text))
  return ""
}


// TODO: Build all the necessary marshal and unmarshal functions

// TODO: This should be marshal actually bc now it imples it generates an Id
// from a string deterministically
func FromString(seed string) (id Id, err error) {
	//i := &Id{}
	//err := i.UnmarshalText([]byte(seed))
	//return *i, err
  return id, err
}
