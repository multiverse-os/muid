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

