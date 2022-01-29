package muid

import (
  "encoding/binary"
  "os"
)

// TODO: Rewrite this, to match our positioning
//func (self Id) Pid() uint16 {
//	return binary.BigEndian.Uint16(self[7:9])
//}


// TODO: Should this be big or little?
func pidBytes() []byte {
  byteBuffer := make([]byte, 2)
  binary.LittleEndian.PutUint16(byteBuffer, uint16(os.Getpid()))
  return byteBuffer
}
