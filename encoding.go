package muid

import (
	"fmt"
	"strings"
  "encoding/base32"
)


func (self Id) Base32()  Id {
  data, _ := base32.StdEncoding.DecodeString(string(self.Bytes()))
  return Id(data)
}

// TODO: This might all be irrelevant and should opt to use standard libraries
// to get these conversions. 

type Encoding uint8

// TODO: Would like to add a portion of the codec library code to this so we can
//       have easy sha3 etc which also work great for id names and keep this
//       library a no depednency library so that it has a larger scope of use
//       among the go community if they decide they like it
const (
	Alphabetic Encoding = iota
	Alphanumeric
	Base32
	Base58
	Numeric
	URLSafe
)

func (self Encoding) Encoding() string {
	switch self {
	case Numeric:
		return "0123456789"
	case Alphabetic:
		return "abcdefghijklmnopqrstuvwxyz"
	case Alphanumeric:
		return fmt.Sprintf("%s%s", Alphabetic.Encoding(), Numeric.Encoding())
	case Base58:
		return fmt.Sprintf("%sABCDEFGHJKLMNPQRSTUVWXYZ%s", Numeric.Encoding(), Alphabetic.Encoding())
	case URLSafe:
		return fmt.Sprintf("%s%s%s_-", Alphabetic.Encoding(), strings.ToUpper(Alphabetic.Encoding()), Numeric.Encoding())
	default: //case Base32:
		return fmt.Sprintf("%sabcdefghijklmnopqrstuv", Numeric.Encoding())
	}
}
