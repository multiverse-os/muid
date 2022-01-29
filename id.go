package muid

import (
  "encoding/base32"
  "encoding/base64"
  //"encoding/hex"
	"fmt"
	"math/rand"
  //"strings"
	"time"
)

type Id []byte

func init() { rand.Seed(time.Now().UTC().UnixNano()) }

func NilId() Id { return Id{} }
func New() Id { return NewWithTime(time.Now()) }

// TODO: Add chainable methods like this for customizing the id 
func (self Id) Prefix(prefix string) Id {
  return self
}

// TODO: For now lets just do full size and we can make .Short() use our
// compressed values
func NewWithTime(timestamp time.Time) Id {
  id := make([]byte, 12)
  copy(id[0:], timestampBytes(timestamp))
  copy(id[4:], pidBytes())
  copy(id[6:], machineIdBytes(3))
  copy(id[9:], randomBytes(2))
  id[11] = simpleChecksumByte(id[:11])
  // NOTE: Id is ready to return here, this is how small we got the primary
  // function, and its far easier to understand, make modification, or
  // customize. Below is development code for completing the alpha version of
  // the library and will all be broken off into their own functions. 
  fmt.Println("id bytes: ", id)
  fmt.Println("string(id bytes): ", string(id))

  encoder := base32.NewEncoding("0123456789abcdefghijklmnopqrstuv").WithPadding(base32.NoPadding)
  base32Id := encoder.EncodeToString(id)
  fmt.Println("custom encoder:", base32Id)


	//encoder := base32.NewEncoder(base32.HexEncoding, os.Stdout)
	//encoder.Write(id)
	//// Must close the encoder when finished to flush any partial blocks.
	//// If you comment out the following line, the last partial block "r"
	//// won't be encoded.
	//encoder.Close()

  // NOTE: This produces a nice 16 character string vs base32 producing 20
  // character
  testBase64 := base64.URLEncoding.EncodeToString(id)
  fmt.Println("test base 64:", testBase64)

  testRawBase64 := base64.RawURLEncoding.EncodeToString(id)
  fmt.Println("test base 64:", testRawBase64)

  // TODO: Then merge all together, then base32 + hex
  //       then its ready to use
  //hexId := hex.EncodeToString(id)
  //fmt.Println("hex version of id: ", hexId)
  //fmt.Println("byte slice version of hex id: ", []byte(hexId))

	//base32Id := base32.StdEncoding.EncodeToString(id[:])
  //
  //compressedBase32Id := strings.Replace(base32Id, "=", "", -1)



  //fmt.Println("base32 version of id: ", base32Id)
  //fmt.Println("base32 version of id: ", compressedBase32Id)
  //fmt.Println("bytes of base32 version of id", []byte(compressedBase32Id))

  //fmt.Println("now we downcase it because duh")

  //idString := strings.ToLower(compressedBase32Id)

  //fmt.Println("finalized id string:", idString)
  //fmt.Println("finalized id string bytes:", []byte(idString))

  //fmt.Println("byte slice version of base32 id: ", []byte(idString))

  //fmt.Println("hexify each version upcase and downcase")

  //compressedHex := hex.EncodeToString([]byte(compressedBase32Id))

  //fmt.Println("compressed hex:", compressedHex)



  // TODO Add ascii58 example
  //b85 := make([]byte, ascii85.MaxEncodedLen(len(t)))
	//n, _, _ := ascii85.Decode(b85, t, true)
  
	return Id(id)
}

func (self Id) IsNil() bool { return (self == nil || len(self) == 0) }

// TODO: Use the checksum if it exists and check not nil and not below minimum size or over maximum
func (self Id) IsValid() bool { return true } 
