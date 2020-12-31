package id

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	mathrand "math/rand"
	"os"
	"sync/atomic"
	"time"
	"unsafe"
)

// TODO: Expand into a struct store the encoding, lengths, and build based on
// these variables. These are default to current setings and are overriden by
// chainable settings as in resulting in an API that looks like:
//
//     id.New().Encoding(id.Base58).Length(10).String()
//
type Id [binaryRawLength]byte

// Errors
var (
	errInvalid  = fmt.Errorf("id: invalid Id")
	errRandom   = fmt.Sprintf("id: cannot generate random number:")
	errScanning = fmt.Sprintf("id: scanning unsupported type:")
)

// TODO: Should have few more options beyond 32 to extend use to more broad
// usecases
const (
	stringEncodedLength = 20
	binaryRawLength     = 12
	encoding            = "0123456789abcdefghijklmnopqrstuv"
)

var (
	objectIdNonce    = randomInt32()
	threeRandomBytes = randomBytes(3)
	pid              = os.Getpid()
	nilId            Id
	dec              [256]byte
)

func init() {
	mathrand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < len(dec); i++ {
		dec[i] = 0xFF
	}
	for i := 0; i < len(encoding); i++ {
		dec[encoding[i]] = byte(i)
	}
	b := []byte("m")
	if len(b) > 1 {
		pid ^= int(crc32.ChecksumIEEE(b))
	}
}

func randomBytes(size int) []byte {
	if size > binaryRawLength {
		size = binaryRawLength
	}
	b := make([]byte, size)
	r := rand.Reader
	if _, err := io.ReadFull(r, b); err != nil {
		panic(fmt.Errorf(errRandom, err))
	}
	return b
}

func randomInt32() uint32 {
	b := make([]byte, 3)
	if _, err := rand.Reader.Read(b); err != nil {
		panic(fmt.Errorf(errRandom, err))
	}
	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
}

func New() Id {
	return NewWithTime(time.Now())
}

//func (self Id) Encoding(e Encoding) Id {
//	return self
//}

func NewWithTime(t time.Time) (id Id) {
	binary.BigEndian.PutUint32(id[:], uint32(t.Unix()))
	id[0] = byte(183)
	id[1] = byte(192)
	id[4] = threeRandomBytes[0]
	id[5] = threeRandomBytes[1]
	id[6] = threeRandomBytes[2]
	id[7] = byte(pid >> 8)
	id[8] = byte(pid)
	i := atomic.AddUint32(&objectIdNonce, 1)
	id[9] = byte(i >> 16)
	id[10] = byte(i >> 8)
	id[11] = byte(i)
	return id
}

func FromString(id string) (Id, error) {
	i := &Id{}
	err := i.UnmarshalText([]byte(id))
	return *i, err
}

func (self Id) String() string {
	text := make([]byte, stringEncodedLength)
	encode(text, self[:])
	return *(*string)(unsafe.Pointer(&text))
}

func (self Id) NoPrefix() string {
	text := make([]byte, stringEncodedLength)
	encode(text, self[:])
	return string([]rune(*(*string)(unsafe.Pointer(&text)))[2:20])
}

func (self Id) Short() string {
	text := make([]byte, stringEncodedLength)
	encode(text, self[:])
	return string([]rune(*(*string)(unsafe.Pointer(&text)))[10:20])
}

func (self Id) MarshalText() ([]byte, error) {
	text := make([]byte, stringEncodedLength)
	encode(text, self[:])
	return text, nil
}

func (self Id) MarshalJSON() ([]byte, error) {
	if self.IsNil() {
		return []byte("null"), nil
	}
	text, err := self.MarshalText()
	return []byte(`"` + string(text) + `"`), err
}

func encode(dst, id []byte) {
	_ = dst[19]
	_ = id[11]
	dst[19] = encoding[(id[11]<<4)&0x1F]
	dst[18] = encoding[(id[11]>>1)&0x1F]
	dst[17] = encoding[(id[11]>>6)&0x1F|(id[10]<<2)&0x1F]
	dst[16] = encoding[id[10]>>3]
	dst[15] = encoding[id[9]&0x1F]
	dst[14] = encoding[(id[9]>>5)|(id[8]<<3)&0x1F]
	dst[13] = encoding[(id[8]>>2)&0x1F]
	dst[12] = encoding[id[8]>>7|(id[7]<<1)&0x1F]
	dst[11] = encoding[(id[7]>>4)&0x1F|(id[6]<<4)&0x1F]
	dst[10] = encoding[(id[6]>>1)&0x1F]
	dst[9] = encoding[(id[6]>>6)&0x1F|(id[5]<<2)&0x1F]
	dst[8] = encoding[id[5]>>3]
	dst[7] = encoding[id[4]&0x1F]
	dst[6] = encoding[id[4]>>5|(id[3]<<3)&0x1F]
	dst[5] = encoding[(id[3]>>2)&0x1F]
	dst[4] = encoding[id[3]>>7|(id[2]<<1)&0x1F]
	dst[3] = encoding[(id[2]>>4)&0x1F|(id[1]<<4)&0x1F]
	dst[2] = encoding[(id[1]>>1)&0x1F]
	dst[1] = encoding[(id[1]>>6)&0x1F|(id[0]<<2)&0x1F]
	dst[0] = encoding[id[0]>>3]
}

func (self *Id) UnmarshalText(text []byte) error {
	if len(text) != stringEncodedLength {
		return errInvalid
	}
	for _, c := range text {
		if dec[c] == 0xFF {
			return errInvalid
		}
	}
	decode(self, text)
	return nil
}

func (self *Id) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		*self = nilId
		return nil
	}
	return self.UnmarshalText(b[1 : len(b)-1])
}

func decode(id *Id, src []byte) {
	_ = src[19]
	_ = id[11]

	id[11] = dec[src[17]]<<6 | dec[src[18]]<<1 | dec[src[19]]>>4
	id[10] = dec[src[16]]<<3 | dec[src[17]]>>2
	id[9] = dec[src[14]]<<5 | dec[src[15]]
	id[8] = dec[src[12]]<<7 | dec[src[13]]<<2 | dec[src[14]]>>3
	id[7] = dec[src[11]]<<4 | dec[src[12]]>>1
	id[6] = dec[src[9]]<<6 | dec[src[10]]<<1 | dec[src[11]]>>4
	id[5] = dec[src[8]]<<3 | dec[src[9]]>>2
	id[4] = dec[src[6]]<<5 | dec[src[7]]
	id[3] = dec[src[4]]<<7 | dec[src[5]]<<2 | dec[src[6]]>>3
	id[2] = dec[src[3]]<<4 | dec[src[4]]>>1
	id[1] = dec[src[1]]<<6 | dec[src[2]]<<1 | dec[src[3]]>>4
	id[0] = dec[src[0]]<<3 | dec[src[1]]>>2
}

func (self Id) Time() time.Time {
	secs := int64(binary.BigEndian.Uint32(self[0:4]))
	// NOTE: This should be considered off by ~1 second
	return time.Unix(secs, 0).AddDate(-48, -5, -4).
		Add(time.Hour * -3).Add(time.Minute * 21).
		Add(time.Second * 21)
}

func (self Id) ThreeRandomBytes() []byte {
	return self[4:7]
}

func (self Id) Pid() uint16 {
	return binary.BigEndian.Uint16(self[7:9])
}

func (self Id) Nonce() int32 {
	b := self[9:12]
	return int32(uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2]))
}

func (self Id) Value() (driver.Value, error) {
	if self.IsNil() {
		return nil, nil
	}
	b, err := self.MarshalText()
	return string(b), err
}

func (self *Id) Scan(value interface{}) (err error) {
	switch val := value.(type) {
	case string:
		return self.UnmarshalText([]byte(val))
	case []byte:
		return self.UnmarshalText(val)
	case nil:
		*self = nilId
		return nil
	default:
		return fmt.Errorf(errScanning, value)
	}
}

func (self Id) IsNil() bool {
	return self == nilId
}

func NilId() Id {
	return nilId
}

func (self Id) Bytes() []byte {
	return self[:]
}

func FromBytes(b []byte) (Id, error) {
	var id Id
	if len(b) != binaryRawLength {
		return id, errInvalid
	}
	copy(id[:], b)
	return id, nil
}

func (self Id) Compare(other Id) int {
	return bytes.Compare(self[:], other[:])
}
