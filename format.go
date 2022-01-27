package muid

import (
  "encoding/hex"
)

func (self Id) Hex() string   { return hex.EncodeToString(id[:]) }
func (self Id) Bytes() []byte { return self[:] }
