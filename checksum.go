package muid

import (
  "encoding/binary"
	"hash/crc32"
  "hash/adler32"
)

func simpleChecksumByte(idBytes []byte) byte {
  var sum uint8
  for _, idByte := range idBytes {
    sum += uint8(idByte)
  }
  return byte(sum / uint8(len(idBytes)))
}

func simpleChecksumValid(idBytes []byte, checksumByte byte) bool {
  return simpleChecksumByte(idBytes) == checksumByte
}

func crc32ChecksumBytes(id Id) []byte {
  byteBuffer := make([]byte, 4)
  binary.BigEndian.PutUint32(byteBuffer, crc32.ChecksumIEEE(id))
  return byteBuffer
}

func crc32ChecksumValid(id Id, checksumBytes []byte) bool {
  return crc32ChecksumBytes(id) == checksumBytes
}

func adler32ChecksumBytes(id Id) []byte {
  byteBuffer := make([]byte, 4)
  binary.BigEndian.PutUint32(byteBuffer, adler32.Checksum(id))
  return byteBuffer 
}

func adler32ChecksumValid(id Id, checksumBytes []byte) bool {
  return adler32ChecksumBytes(id) == checksumBytes
}
