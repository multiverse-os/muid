package muid

import (
  "bytes"
  "encoding/base32"
  "encoding/base64"
  "encoding/hex"
	"math/rand"
  "sort"
	"time"
)

type Id []byte

func init() { rand.Seed(time.Now().UTC().UnixNano()) }

// Generation
func NilId() Id { return Id{} }

func Generate() Id { return AtTime(time.Now()) }

func Deterministic(seed string, length int) Id {
  if length < len(seed) {
    length = len(seed)
  }
  id := make([]byte, length)
  if len(seed) < length {
    randomByteLength := length - len(seed)
    copy(id[len(seed):], hashBytes(Shake256, []byte(seed), randomByteLength))
  }
  return Id(id)
}

func AtTime(timestamp time.Time) Id {
  id := make([]byte, 12)
  copy(id[0:], timestampBytes(timestamp))
  copy(id[4:], pidBytes())
  copy(id[6:], machineIdBytes(3))
  copy(id[9:], randomBytes(2))
  id[11] = checksumByte(id[:11])
  return Id(id)
}

// Validations
func (self Id) IsNil() bool { return (self == nil || len(self) == 0) }
func (self Id) ChecksumValid() bool { return checksumValid(self[:11], self[11]) }
func (self Id) IsValid() bool { return !self.IsNil() && self.ChecksumValid() } 

// Sorting
type sorter []Id

func (self Id) Compare(other Id) int { return bytes.Compare(self[:], other[:]) }

func Sort(ids []Id) { sort.Sort(sorter(ids)) }

func (self sorter) Len() int { return len(self) }
func (self sorter) Less(i, j int) bool { return self[i].Compare(self[j]) < 0 }
func (self sorter) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

// Format & Encoding
func (self Id) String() string { return string(self[:]) }
func (self Id) Bytes()  []byte { return self[:] }
func (self Id) Hex()    string { return hex.EncodeToString(self[:]) }

func (self Id) Base32() Id {
  encoder := base32.NewEncoding("0123456789abcdefghijklmnopqrstuv").WithPadding(base32.NoPadding)
  base32Id := encoder.EncodeToString(self)
  return Id(base32Id)
}

func (self Id) Base64() Id {
  base64Id := base64.URLEncoding.EncodeToString(self)
  return Id(base64Id)
}

//func (self Id) Base58() Id {
//  // TODO Add ascii58 example
//  //b85 := make([]byte, ascii85.MaxEncodedLen(len(t)))
//	//n, _, _ := ascii85.Decode(b85, t, true)
//  return self
//}

// TODO: JSON will help reach feature parity with `bsonid`
//func (self *Id) UnmarshalJSON(b []byte) error {
//	s := string(b)
//	if s == "null" {
//		*self = nilId
//		return nil
//	}
//	return self.UnmarshalText(b[1 : len(b)-1])
//}

//func (self Id) MarshalJSON() ([]byte, error) {
//	if self.IsNil() {
//		return []byte("null"), nil
//	}
//	text, err := self.MarshalText()
//	return []byte(`"` + string(text) + `"`), err
//}

// Marshalling 
// TODO: Add code to marshal bytes, string, hex, base32, base64 versions back
//       into an Id object
