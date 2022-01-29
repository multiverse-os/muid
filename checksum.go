package muid

import (
  "bytes"
  "encoding/binary"
	"hash/crc32"
  "hash/adler32"
)

func checksumByte(idBytes []byte) byte {
  var sum uint8
  for _, idByte := range idBytes {
    sum += uint8(idByte)
  }
  return byte(sum / uint8(len(idBytes)))
}

func checksumValid(idBytes []byte, checkByte byte) bool {
  return checksumByte(idBytes) == checkByte
}

func crc32ChecksumBytes(id Id) []byte {
  byteBuffer := make([]byte, 4)
  binary.BigEndian.PutUint32(byteBuffer, crc32.ChecksumIEEE(id))
  return byteBuffer
}

func crc32ChecksumValid(id Id, checksumBytes []byte) bool {
  return bytes.Compare(crc32ChecksumBytes(id), checksumBytes) == 0
}

func adler32ChecksumBytes(id Id) []byte {
  byteBuffer := make([]byte, 4)
  binary.BigEndian.PutUint32(byteBuffer, adler32.Checksum(id))
  return byteBuffer 
}

func adler32ChecksumValid(id Id, checksumBytes []byte) bool {
  return bytes.Compare(adler32ChecksumBytes(id), checksumBytes) == 0
}
