package muid

import (
  "bytes"
  "encoding/binary"
	"hash/crc32"
  "hash/adler32"
)

//type ChecksumType uint8
//
//const (
//  Undefined ChecksumType = iota
//  MD5
//  SHA1
//  SHA224
//  SHA256
//  SHA384
//  SHA512
//  SHA3_224
//  SHA3_256
//  SHA3_384
//  SHA3_512
//)

//func checksum(checksumType ChecksumType, data []byte) (checksum string, err error) {
//	// default
//	var hasher hash.Hash
//	switch checksumType {
//	//case MD5:
//	//	hasher = md5.New()
//	//case SHA1:
//	//	hasher = sha1.New()
//	//case SHA224:
//	//	hasher = sha256.New224()
//	//case SHA256:
//	//	hasher = sha256.New()
//	//case SHA384:
//	//	hasher = sha512.New384()
//	//case SHA512:
//	//	hasher = sha512.New()
//	case SHA3_224:
//		hasher = sha3.New224()
//	case SHA3_256:
//		hasher = sha3.New256()
//	case SHA3_384:
//		hasher = sha3.New384()
//	case SHA3_512:
//		hasher = sha3.New512()
//	default:
//		msg := "Invalid algorithm parameter passed go Checksum: %s"
//		return checksum, fmt.Errorf(msg, algorithm)
//	}
//	hasher.Write(data)
//	str := hex.EncodeToString(hasher.Sum(nil))
//	return str, nil
//}

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
