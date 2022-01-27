package muid

import (
  "io"
  "crypto/rand"
)

// TODO: Length should be variable
const (
  binaryRawLength     = 12
)

func RandomBytes(size int) []byte {
	if binaryRawLength < size {
		size = binaryRawLength
	}
	byteBuffer := make([]byte, size)
	randReader := rand.Reader
	if _, err := io.ReadFull(randReader, byteBuffer); err != nil {
		panic(err)
	}
	return byteBuffer
}

func RandomInt32() uint32 {
	byteBuffer := make([]byte, 3)
	if _, err := rand.Reader.Read(byteBuffer); err != nil {
		panic(err)
	}
	return uint32(byteBuffer[0])<<16 | uint32(byteBuffer[1])<<8 | uint32(byteBuffer[2])
}
