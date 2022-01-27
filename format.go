package muid

import (
  "encoding/hex"
)

func (self Id) Hex() string   { return hex.EncodeToString(self[:]) }
func (self Id) Bytes() []byte { return self[:] }
