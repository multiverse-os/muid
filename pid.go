package muid

import (
  "encoding/binary"
  "os"
)

// TODO: Rewrite this, to match our positioning
//func (self Id) Pid() uint16 {
//	return binary.BigEndian.Uint16(self[7:9])
//}

func pidBytes() []byte {
  byteBuffer := make([]byte, 2)
  binary.BigEndian.PutUint16(byteBuffer, uint16(os.Getpid()))
  return byteBuffer
}
