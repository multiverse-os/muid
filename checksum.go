package muid

import (
  "encoding/binary"
	"hash/crc32"
  "hash/adler32"
)

// NOTE: Because we use time, the values can never be all zero,
//       so this method will work reliably. 
func simpleChecksumByte(idBytes []byte) byte {
  var sum uint8
  for _, idByte := range idBytes {
    sum += uint8(idByte)
  }
  return byte(sum / uint8(len(idBytes)))
}

func crc32ChecksumBytes(id Id) []byte {
  byteBuffer := make([]byte, 4)
  binary.BigEndian.PutUint32(byteBuffer, crc32.ChecksumIEEE(id))
  return byteBuffer
}

func adler32ChecksumBytes(id Id) []byte {
  byteBuffer := make([]byte, 4)
  binary.BigEndian.PutUint32(byteBuffer, adler32.Checksum(id))
  return byteBuffer 
}
