package muid

import (
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

// TODO: Was already leaning towards something like this but for reference xid
// uses: 
// Base32 hex encoded by default (16 bytes storage when transported as printable string)

// TODO: xid is adds so much fucking code just to get a downcased version of
// base32. When you can just do the normal base32 and downcase it (and upcase
// on reverse operations)... o.O

type Id []byte

func init() { rand.Seed(time.Now().UTC().UnixNano()) }

func NilId() Id { return Id{} }
func New() Id { return NewWithTime(time.Now()) }



// TODO: Add chainable methods like this for customizing the id 
func (self Id) Prefix(prefix string) Id {
  return self
}

func NewWithTime(timestamp time.Time) Id {
  id := make([]byte, 6, 64)

  copy(id[0:4], timestampBytes(timestamp))
  fmt.Println("id bytes: ", id)
  copy(id[5:6], pidBytes())
  fmt.Println("id bytes: ", id)
  // TODO copy(id[7:8], machineBytes())
  // TODO copy(id[9:10], checksumBytes())
  // TODO copy(id[8:], randomBytes(length))

  machineBytes := machineIdBytes(3)
  fmt.Println("machine bytes (3): ", machineBytes)

  // TODO: Then merge all together, then base32 + hex
  //       then its ready to use
  hexId := hex.EncodeToString(id)
  fmt.Println("hex version of id: ", hexId)
  fmt.Println("byte slice version of hex id: ", []byte(hexId))

	base32Id := base32.StdEncoding.EncodeToString(id[:])
  fmt.Println("base32 version of id: ", base32Id)
  fmt.Println("byte slice version of base32 id: ", []byte(base32Id))

  // TODO Add ascii58 example
  //b85 := make([]byte, ascii85.MaxEncodedLen(len(t)))
	//n, _, _ := ascii85.Decode(b85, t, true)
  
	return Id(id)
}

func (self Id) IsNil() bool { return (self == nil || len(self) == 0) }
// TODO: Use the checksum if it exists and check not nil and not below minimum size or over maximum
func (self Id) IsValid() bool { return true } 


// NOTE: Removing all excess code, we have divered so far that their misnamed
// marshal functions are unlikely to relevant and at the very least could be
// better written so we will glance over xid and ensure we did not lose any
// functionality. 
//
// Our goal is less code, cleaner, easier-to-understand, less memory, more
// functionality, and smaller or highly customizable ids for the widest
// number of usecases. 
