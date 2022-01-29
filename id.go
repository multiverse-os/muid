package muid

import (
	"encoding/binary"
  "encoding/base32"
  "encoding/hex"
	"fmt"
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

func (self Id) IsNil() bool { return (self == nil || len(self) == 0) }
// TODO: Use the checksum if it exists and check not nil and not below minimum size or over maximum
func (self Id) IsValid() bool { return true } 

func NilId() Id { return Id{} }
func New() Id { return NewWithTime(time.Now()) }

// TODO: Add chainable methods like this for customizing the id 
func (self Id) Prefix(prefix string) Id {
  return self
}


func NewWithTime(timestamp time.Time) Id {
  var id []byte 
  id = make([]byte, 8, 64)

  copy(id[0:4], timestampBytes())
  copy(id[5:6], pidBytes())
  // TODO copy(id[7:8], machineBytes())
  // TODO copy(id[8:], randomBytes(length))
  // TODO copy(



  // TODO: Get the machine id and convert it to binary (can we get 1 byte? maybe
  // by using the checksum and pulling first byte or hashing and first and
  // last?)

  // TODO: Get some number of random bytes - this will help either make it work
  // if ran in parallel on a massive amount of amchine or allow for a shorter id 
  // format depending on size

  // TODO: Then merge all together, then base32 + hex
  //       then its ready to use


  // TODO: The output string will be based on a chainable but these need to go
  // into encoding so we can simplify the chainable by calling internal
  // functions 
  hexId := hex.EncodeToString(id)
  fmt.Println("hex version of id: ", hexId)
  fmt.Println("byte slice version of hex id: ", []byte(hexId))

  // TODO Add ascii58 example
  //b85 := make([]byte, ascii85.MaxEncodedLen(len(t)))
	//n, _, _ := ascii85.Decode(b85, t, true)
  
	base32Id := base32.StdEncoding.EncodeToString(id[:])
  fmt.Println("base32 version of id: ", base32Id)
  fmt.Println("byte slice version of base32 id: ", []byte(base32Id))

  // TODO: The resulting id MUST
  //         * be easily converted to a string that is base32 or base58 or base64
  //           or at least URL safe
  //         * must be able to sort using that string
  //         * easily convert that string back into the id object 
  //         * contain some sort of checksum to valdiate if an id could
  //           be valid or not


  // TODO: Now put in the random bits based on the length

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

// NOTE: Removing all excess code, we have divered so far that their misnamed
// marshal functions are unlikely to relevant and at the very least could be
// better written so we will glance over xid and ensure we did not lose any
// functionality. 
//
// Our goal is less code, cleaner, easier-to-understand, less memory, more
// functionality, and smaller or highly customizable ids for the widest
// number of usecases. 
