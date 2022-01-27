package muid

import (
	//"database/sql/driver"
	"encoding/binary"
  "encoding/base32"
  "encoding/hex"
	"fmt"
	"hash/crc32"
  "hash/adler32"
	//"io"
	"math/rand"
	"os"
	//"sync/atomic"
	"time"
	//"unsafe"
)

// TODO: This is a combination of xid and a few other id projects. I'm trying to
//       refactor the code and merge in more featuers so it will support
//       bsonid and other missing functionality. because the included bsonid 
//       library is written well in some ways and insane in others. 
//       so i want to clean it up and have an actually well written, easy to
//       userderstand, easy to modify, improve, update, and use. 
//       because yes, in its current state its nearly none of that.
//       but it does function.

// TODO: Expand into a struct store the encoding, lengths, and build based on
// these variables. These are default to current setings and are overriden by
// chainable settings as in resulting in an API that looks like:
//
//     id.New().Encoding(id.Base58).Length(10).String()
//
//  ewww
//  more like
// 
//     id.Base58().Length(10).String() 
//
// for reals right?

// TODO: Used by [12]byte but the goal is to move to a variable length
//       Id concept
type Id []byte

// TODO: Create a Length() function that will alter the byte array so that the
// legnth is minimum 

// Errors

// TODO: Should have few more options beyond 32 to extend use to more broad
// usecases. This also should be using the newly build encoding.go file rather
// than this encodnig const string. 
// Id rather not use this memory, id rather just create the object early on and
// fill in these variables so the info remains in the binary not the memory :O 

//const (
//	stringEncodedLength = 20
//  // TODO: Built out the ecoding.go file to replace the need for this and expand
//  // beyond base32 (which this is and wasnt properly labled such) 
//)

// I kinda hate this, if we keep it like this nd not just assinged in the
// function or init then we use memory and binary space instead of just binary
// space 
//var (
  // a real nonce needs to increment up always
	// objectIdNonce    = randomInt32()
	// threeRandomBytes = randomBytes(3)
  // TODO: One thing xid did really really well is this. Using the pid as the
  // unique machine random seed. It is a brilliant solution because it gives a
  // great seed that is typically protected regularly changing and cost very
  // little overhead. 
	//pid              = os.Getpid()
  // TODO: This is totally uncessary 
	//nilId            Id
  // TODO: Really? REally? REALLY? we just need to have this space in the memory
  // locked BECAUSE!
	//   dec              [256]byte
//)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	//for i := 0; i < 256; i++ {
	//	dec[i] = 0xFF
	//}
	//for i := 0; i < len(encoding); i++ {
	//	dec[encoding[i]] = byte(i)
	//}
	//b := []byte("m")
	//if len(b) > 1 {
	//	pid ^= int(crc32.ChecksumIEEE(b))
	//}
}


func New() Id {
	return NewWithTime(time.Now())
}

//func (self Id) Encoding(e Encoding) Id {
//	return self
//}

// TODO: This is it, everything else at the moment is auxillary or doesnt work.
//       1) So comment everything else out execpt those specificlaly using this
//          system or functions needed by this. 
//       2) refactor this function and remove the excess functions, refer to the 
//          bson id version included for hints on how to do this
//       3) then reimplement each of the commented functions from step 1, 
//          or decide if they can be removed. Change the function names to be
//          more sensical and in-line with the terms how they are used by most
//          people because some of the function names pulled from xid are not
//          all correct
//       4) merge in the bson id library so that it can all be done from a
//          single one
//       5) Support noncing (sortability), deterministic, ultra-random (use
//          time, pid and random) 
//       6) Have functions that can take either the bytes or the string
//          version of the id and convert it back into an Id object
//       7) Short() and Prefix and similar utility functions should
//          create a new Id object
//       8) Have output as either bytes or string
//       9) Easy variable length limited but a hard lower limit for security
//      
func NewWithTime(timestamp time.Time) Id {
  // So the most minimal version will be 
  // Time  + Pid (Machine Random) + Regular Random + Checksum
  //                                  (giving us our variable length)

  // Time is already a nonce, it always increments and gives us ability 
  // to sort
  // Making the minimum length 
  //   4 + 3 + X + 4 (< 12)

  // What is definitely missing is a built in checksum that can be divided
  // off and used to verify the rest is correct ensuring that the id is
  // valid. but it would be required to be present at end or start. 
  // technically it could be in the middle but you would ahve to do a 
  // weird process of extracting it and combining the first and half portions
  // then checking it against the checksum 
  // TODO: Just use the id object to store the data instead of creatnig
  //       a possibly uneccessary buffer


  // NOTE: This is what we are aiming to compete with
  // UUID is 16 bytes
  // xid which is heavily inspiring this library is
  //   4 (time) + 3 (machine id) + 2 (process id) + 3 (nonce)
  // xid is 12 bytes

  // NOTE: We theoritically could drop time down to 2 by subtracting a fixed
  // number of years from the unix time and adding them back. Not going into
  // detail but if one knows how unix time works they will understand how this
  // would make sense to reliably drop the 

  // muid is aiming to shirk the machine id and use the time as a nonce
  // to allow for more random bits and a checksum


  // IMPORTANT NOTE ******************************************************
  // TODO (shrink the bitsize of the time) - use time.Year * 30 to get get the
  // seconds instead of hard coding the value for cleaner code.
  // 4 (time) + 2 (process id) + 2 (machine id) + (1-6)(random) + (2-4)(checksum)
  //   4 time + 2 process + 2 checksum (first and last) 8
  //  even with 3 byte process id we can get and using the slim checksum
  //  we can get 3 bytes of random data which is important
  //  otherwise we can also use machine id and get 1 byte of random

  // 2 time + 2 process + 2 machine + 2 random = 8 minimum.. 4 checksum gives us
  // 12 and 2 gives us the smaller 10 byte version
  // **Thisis the most important part of this document currently ********


  // TODO: Was already leaning towards something like this but for reference xid
  // uses: 
  // Base32 hex encoded by default (16 bytes storage when transported as printable string)


  // TODO: xid is adds so much fucking code just to get a downcased version of
  // base32. When you can just do the normal base32 and downcase it (and upcase
  // on reverse operations)... o.O

  var id []byte 
  id = make([]byte, 8, 64)

  var byteBuffer []byte

  pid              := os.Getpid()

  fmt.Println("byte buffer: %v", byteBuffer)

  // TODO: Move to separate prefix method so this is optional
  prefix := []byte("mv")
  copy(id[:], prefix)

  byteBuffer = make([]byte, 4)
	binary.BigEndian.PutUint32(byteBuffer, uint32(timestamp.Unix()))
  fmt.Println("byte buffer: %v", byteBuffer)
  copy(id[2:], byteBuffer)

  // TODO: This version will give us 2 bytes of pid

  // NOTE: By shifting the pid to the right by 8 bytes and by doing this we
  //       but LOL just use the fucking go fuckions. look how I did it below
  //       that is actually readable by most programmers and just as effective
  //fmt.Println("pid >> 8: ", (pid >> 8))
  //fmt.Println("byte(pid >> 8): ", byte(pid >> 8))
  //id[5] = byte(pid >> 8)

	//id[6] = byte(pid)
  //fmt.Println("pid: ", pid)
  //fmt.Println("byte(pid): ", byte(pid))

  //binary.BigEndian.Uint16(id[7:9])


  byteBuffer = make([]byte, 2)
  binary.LittleEndian.PutUint16(byteBuffer, uint16(pid))
  fmt.Println("byte buffer: %v", byteBuffer)
  copy(id[6:], byteBuffer)
  //id = append(id, byteBuffer...)

  // TODO: Get the machine id and convert it to binary  (only need 2 bits) 

  // TODO: Get some number of random bytes

  // TODO: Add checksum

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

  //oiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii so easy!

	//binary.BigEndian.PutUint32(id[:], uint32(t.Unix()))
	//id[0] = byte(183)
	//id[1] = byte(192)
	//id[4] = threeRandomBytes[0]
	//id[5] = threeRandomBytes[1]
	//id[6] = threeRandomBytes[2]
	//id[7] = byte(pid >> 8)
	//id[8] = byte(pid)
	//i := atomic.AddUint32(&objectIdNonce, 1)
	//id[9] = byte(i >> 16)
	//id[10] = byte(i >> 8)
	//id[11] = byte(i)
	return Id(id)
}

func (self Id) Pid() uint16 {
	return binary.BigEndian.Uint16(self[7:9])
}


//func readPlatformMachineID() (string, error) {
//	b, err := ioutil.ReadFile("/etc/machine-id")
//	if err != nil || len(b) == 0 {
//		b, err = ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
//	}
//	return string(b), err
//}

//func readMachineID() []byte {
//	id := make([]byte, 3)
//	hid, err := readPlatformMachineID()
//	if err != nil || len(hid) == 0 {
//		hid, err = os.Hostname()
//	}
//	if err == nil && len(hid) != 0 {
//		hw := md5.New()
//		hw.Write([]byte(hid))
//		copy(id, hw.Sum(nil))
//	} else {
//		// Fallback to rand number if machine id can't be gathered
//		if _, randErr := rand.Reader.Read(id); randErr != nil {
//			panic(fmt.Errorf("xid: cannot get hostname nor generate a random number: %v; %v", err, randErr))
//		}
//	}
//	return id
//}


func (self Id) Timestamp() time.Time {
	unixTime := binary.BigEndian.Uint32(self[0:4])
	return time.Unix(int64(unixTime), 0).UTC()
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

func (self Id) String() string {
	//text := make([]byte, stringEncodedLength)
	//encode(text, self[:])
	//return *(*string)(unsafe.Pointer(&text))
  return ""
}

//func (self Id) NoPrefix() string {
//	text := make([]byte, stringEncodedLength)
//	encode(text, self[:])
//	return string([]rune(*(*string)(unsafe.Pointer(&text)))[2:20])
//}
//
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
//func (self Id) MarshalJSON() ([]byte, error) {
//	if self.IsNil() {
//		return []byte("null"), nil
//	}
//	text, err := self.MarshalText()
//	return []byte(`"` + string(text) + `"`), err
//}

// TODO: I hate this function, this can be done in a way that is easy
//       to understand and modify because most developer encoutering
//       this will start crying because they dont understand what 
//       is happening. And Go provides us with all the tools (esp 1.9)
//       to do this exactly functionality but in such a way that is
//       'elegant' syntax
//        OH BIG TALK JERK! WE#LL LETS SEE IT!
//       ok :
//            binary.BigEndian.PutUint32(b[0:4], uint32(timestamp.Unix()))
//       OH THAT IS NICE
//func encode(dst, id []byte) {
//	_ = dst[19]
//	_ = id[11]
//	dst[19] = encoding[(id[11]<<4)&0x1F]
//	dst[18] = encoding[(id[11]>>1)&0x1F]
//	dst[17] = encoding[(id[11]>>6)&0x1F|(id[10]<<2)&0x1F]
//	dst[16] = encoding[id[10]>>3]
//	dst[15] = encoding[id[9]&0x1F]
//	dst[14] = encoding[(id[9]>>5)|(id[8]<<3)&0x1F]
//	dst[13] = encoding[(id[8]>>2)&0x1F]
//	dst[12] = encoding[id[8]>>7|(id[7]<<1)&0x1F]
//	dst[11] = encoding[(id[7]>>4)&0x1F|(id[6]<<4)&0x1F]
//	dst[10] = encoding[(id[6]>>1)&0x1F]
//	dst[9] = encoding[(id[6]>>6)&0x1F|(id[5]<<2)&0x1F]
//	dst[8] = encoding[id[5]>>3]
//	dst[7] = encoding[id[4]&0x1F]
//	dst[6] = encoding[id[4]>>5|(id[3]<<3)&0x1F]
//	dst[5] = encoding[(id[3]>>2)&0x1F]
//	dst[4] = encoding[id[3]>>7|(id[2]<<1)&0x1F]
//	dst[3] = encoding[(id[2]>>4)&0x1F|(id[1]<<4)&0x1F]
//	dst[2] = encoding[(id[1]>>1)&0x1F]
//	dst[1] = encoding[(id[1]>>6)&0x1F|(id[0]<<2)&0x1F]
//	dst[0] = encoding[id[0]>>3]
//}

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
//func (self *Id) UnmarshalJSON(b []byte) error {
//	s := string(b)
//	if s == "null" {
//		*self = nilId
//		return nil
//	}
//	return self.UnmarshalText(b[1 : len(b)-1])
//}

//func decode(id *Id, src []byte) {
//	_ = src[19]
//	_ = id[11]
//
//	id[11] = dec[src[17]]<<6 | dec[src[18]]<<1 | dec[src[19]]>>4
//	id[10] = dec[src[16]]<<3 | dec[src[17]]>>2
//	id[9] = dec[src[14]]<<5 | dec[src[15]]
//	id[8] = dec[src[12]]<<7 | dec[src[13]]<<2 | dec[src[14]]>>3
//	id[7] = dec[src[11]]<<4 | dec[src[12]]>>1
//	id[6] = dec[src[9]]<<6 | dec[src[10]]<<1 | dec[src[11]]>>4
//	id[5] = dec[src[8]]<<3 | dec[src[9]]>>2
//	id[4] = dec[src[6]]<<5 | dec[src[7]]
//	id[3] = dec[src[4]]<<7 | dec[src[5]]<<2 | dec[src[6]]>>3
//	id[2] = dec[src[3]]<<4 | dec[src[4]]>>1
//	id[1] = dec[src[1]]<<6 | dec[src[2]]<<1 | dec[src[3]]>>4
//	id[0] = dec[src[0]]<<3 | dec[src[1]]>>2
//}

//func (self Id) ThreeRandomBytes() []byte {
//	return self[4:7]
//}
//
//func (self Id) Pid() uint16 {
//	return binary.BigEndian.Uint16(self[7:9])
//}

//func (self Id) Nonce() int32 {
//	b := self[9:12]
//	return int32(uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2]))
//}
//
//func (self Id) Value() (driver.Value, error) {
//	if self.IsNil() {
//		return nil, nil
//	}
//	b, err := self.MarshalText()
//	return string(b), err
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
//
//func (self Id) IsNil() bool {
//	return self == nilId
//}
//
//func NilId() Id {
//	return nilId
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

