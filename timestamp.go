package muid

import (
  "time"
  "encoding/binary"
)

// NOTE: Amount of time is not really important, just consistency 
//       (value is based on 8760 hours per year)
func fourtyYears() time.Duration {
  return (time.Hour * 350400)
}

func compressTimestamp(timestamp time.Time) uint16 {
  return uint16(uint32(timestamp.Unix()) - uint32(fourtyYears().Seconds()))
}

func uncompressTimestamp(compressedTimestamp uint16) time.Time {
  return time.Unix(
    int64(uint32(compressedTimestamp) + uint32(fourtyYears().Seconds())),
    0,
  ).UTC()
}

func compressedTimestampBytes(timestamp time.Time) []byte {
  byteBuffer := make([]byte, 2)
  binary.BigEndian.PutUint16(byteBuffer, compressTimestamp(timestamp))
  return byteBuffer
}

func timestampBytes(timestamp time.Time) []byte {
  byteBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(byteBuffer, uint32(timestamp.Unix()))
  return byteBuffer
}

// TODO: Need to detect if it fails and retry with compressed version
func (self Id) Timestamp() time.Time {
	unixTime := binary.BigEndian.Uint32(self[0:4])
	return time.Unix(int64(unixTime), 0).UTC()
}
