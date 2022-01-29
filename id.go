package muid

import (
	"encoding/binary"
  "encoding/base32"
  "encoding/hex"
	"fmt"
	"hash/crc32"
  "hash/adler32"
	"math/rand"
	"time"
)

// TODO: This is a combination of xid and a few other id projects. I'm trying to
//       refactor the code and merge in more featuers so it will support
//       bsonid and other missing functionality. because the included bsonid 
//       library is written well in some ways and insane in others. 
//       so i want to clean it up and have an actually well written, easy to
//       userderstand, easy to modify, improve, update, and use. 
//       because yes, in its current state its nearly none of that.
//       but it does function.

//       another important goal is to also being able to produce a wide variety
//       of similar ids, like the end result of this library should also be able
//       to produce mongoDB bson style ids (beause the current go library for it
//       its awful)

//     id.Base58().Length(10).String() 
//
// NOTE: This is what we are aiming to compete with
// UUID is 16 bytes
// xid is 12 bytes
// xid which is heavily inspiring this library is
//   4 (time) + 3 (machine id) + 2 (process id) + 3 (nonce)

// muid 8..12 bytes depending on developer requirements
// 2 time + 2 process + 2 machine + 2 random = 8 minimum.. 4 checksum gives us
// 12 and 2 gives us the smaller 10 byte version

// TODO: Was already leaning towards something like this but for reference xid
// uses: 
// Base32 hex encoded by default (16 bytes storage when transported as printable string)

// TODO: xid is adds so much fucking code just to get a downcased version of
// base32. When you can just do the normal base32 and downcase it (and upcase
// on reverse operations)... o.O


type Id []byte

func init() { rand.Seed(time.Now().UTC().UnixNano()) }

func (self Id) IsNil() bool { return self == nil }

func NilId() Id { return Id{} }
func New() Id { return NewWithTime(time.Now()) }

func NewWithTime(timestamp time.Time) Id {
  var id []byte 
  id = make([]byte, 8, 64)

  var byteBuffer []byte

  fmt.Println("byte buffer: %v", byteBuffer)

  // TODO: Move to separate prefix method so this is optional
  prefix := []byte("mv")
  copy(id[:], prefix)

  byteBuffer = make([]byte, 4)

  // NOTE: This originally bigEndian which is why its like that now but we
  // experimented with using littleEndian  which is TECHNICALLY PREFERRED since
  // it would put the nonce at the start and therefore make it much simler to
  // sort by





  // NOTE: Doing it little endian gives us 2 bytes vs big endian gives us 4
  //pid              := os.Getpid()
  //byteBuffer = make([]byte, 2)
  //binary.LittleEndian.PutUint16(byteBuffer, uint16(pid))
  //fmt.Println("byte buffer: %v", byteBuffer)
  // NOTE: All above was replaced with this single line, LIKE WOW o.O im not
  // even a good programmer
  copy(id[6:], pidBytes())



  // TODO: Get the machine id and convert it to binary (can we get 1 byte? maybe
  // by using the checksum and pulling first byte or hashing and first and
  // last?)

  // TODO: Get some number of random bytes - this will help either make it work
  // if ran in parallel on a massive amount of amchine or allow for a shorter id 
  // format depending on size

  // TODO: Then merge all together, then base32 + hex
  //       then its ready to use


   
  fmt.Println("byte slice version of id: ", id)
  fmt.Println("string version of id: ", string(id))

  hexId := hex.EncodeToString(id)
  fmt.Println("hex version of id: ", hexId)
  fmt.Println("byte slice version of hex id: ", []byte(hexId))

  // TODO Add ascii58 example
  //b85 := make([]byte, ascii85.MaxEncodedLen(len(t)))
	//n, _, _ := ascii85.Decode(b85, t, true)
  
	base32Id := base32.StdEncoding.EncodeToString(id[:])
  fmt.Println("base32 version of id: ", base32Id)
  fmt.Println("byte slice version of base32 id: ", []byte(base32Id))

  // NOTE: We could just use 2 bytes and do the checksum and only check against
  // first and last. In this way we could use sha3 and do something similar

  // TODO: Add the checksum with a method to allow developer to pick which they
  // prefer OR more importantly if they want to include it. 
  var crc32Id []byte
  crc32Id = make([]byte, 4)
  binary.BigEndian.PutUint32(crc32Id, crc32.ChecksumIEEE(id))
  fmt.Println("crc32 version of id: ", crc32Id)
  fmt.Println("string version of crc32 id: ", string(crc32Id))

  var adler32Id []byte
  adler32Id = make([]byte, 4)
  binary.BigEndian.PutUint32(adler32Id, adler32.Checksum(id))
  fmt.Println("adler32 version of id: ", adler32Id)
  fmt.Println("string version of adler32 id: ", string(adler32Id))

  // TODO: The resulting id MUST
  //         * be easily converted to a string that is base32 or base58 or base64
  //           or at least URL safe
  //         * must be able to sort using that string
  //         * easily convert that string back into the id object 
  //         * contain some sort of checksum to valdiate if an id could
  //           be valid or not


  // TODO: Now put in the random bits based on the length

  // TODO: Then generate the checksum of that and attach it somewhere

  // TODO: Then output the newly generated id

	return Id(id)
}

// TODO: This should be marshal actually bc now it imples it generates an Id
// from a string deterministically
func FromString(seed string) (id Id, err error) {
	//i := &Id{}
	//err := i.UnmarshalText([]byte(seed))
	//return *i, err
  return id, err
}

//func MarshalId(idString string) (id Id, err error) {
//  return id, err
//}

//func (self Id) Short() string {
//	text := make([]byte, stringEncodedLength)
//	encode(text, self[:])
//	return string([]rune(*(*string)(unsafe.Pointer(&text)))[10:20])
//}

// TODO: Marshal text should not be a method of the object, it should take in
// bytes and return the Id. So either this is mis-named or it is improperly
// implemented
//func (self Id) MarshalText() ([]byte, error) {
//	text := make([]byte, stringEncodedLength)
//	encode(text, self[:])
//	return text, nil
//}
//

//func (self *Id) UnmarshalText(text []byte) error {
//	if len(text) != stringEncodedLength {
//		return errInvalid
//	}
//	for _, c := range text {
//		if dec[c] == 0xFF {
//			return errInvalid
//		}
//	}
//	decode(self, text)
//	return nil
//}
//

//func (self Id) ThreeRandomBytes() []byte {
//	return self[4:7]
//}
//

//func (self *Id) Scan(value interface{}) (err error) {
//	switch val := value.(type) {
//	case string:
//		return self.UnmarshalText([]byte(val))
//	case []byte:
//		return self.UnmarshalText(val)
//	case nil:
//		*self = nilId
//		return nil
//	default:
//		return fmt.Errorf(errScanning, value)
//	}
//}


//func FromBytes(b []byte) (Id, error) {
//	var id Id
//	if len(b) != binaryRawLength {
//		return id, errInvalid
//	}
//	copy(id[:], b)
//	return id, nil
//}
//




