package muid

import (
  "encoding/binary"
	"hash/crc32"
  "hash/adler32"
)

// TODO: It will be necessary to add in the code for a compressed
//       version of the checksum, like using first and last 
//       In addition we will need to reverse functions for 
//       taking a Id and validating the checksum. Support should
//       exist for both short and full versions of checksums
//       to give developers more options and expand our number
//       of potential use cases. 

func crc32ChecksumBytes(id Id) []byte {
  byteBuffer = make([]byte, 4)
  binary.BigEndian.PutUint32(byteBuffer, crc32.ChecksumIEEE(id))
  return byteBuffer
}

func adler32ChecksumBytes(id Id) []byte {
  byteBuffer = make([]byte, 4)
  binary.BigEndian.PutUint32(byteBuffer, adler32.Checksum(id))
  return byteBuffer 
}
