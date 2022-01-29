package muid

import (
  "encoding/binary"
  "os"
)

func (self Id) Pid() uint16 { return binary.BigEndian.Uint16(self[4:5]) }

func pidBytes() []byte {
  byteBuffer := make([]byte, 2)
  binary.BigEndian.PutUint16(byteBuffer, uint16(os.Getpid()))
  return byteBuffer
}
