package muid

import (
  "time"
  "encoding/binary"
)

// TODO: Need to add a method of Id to optionally use either full or compressed
//       version of the timestamp to support a wider number of use-cases 

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

// TODO: This appears to work well in poc form, we should flip the endian and
//       experiment using this in the id over non-compressed time, but
//       in the end this should be developer option
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

// TODO: This has to be updated using our compress and uncompress timestamp
// functions
//func (self Id) Timestamp() time.Time {
//	unixTime := binary.BigEndian.Uint32(self[0:4])
//	return time.Unix(int64(unixTime), 0).UTC()
//}
