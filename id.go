package muid

import (
	"math/rand"
	"time"
)

type Id []byte

func init() { rand.Seed(time.Now().UTC().UnixNano()) }

func NilId() Id { return Id{} }

func Generate() Id { return AtTime(time.Now()) }

func AtTime(timestamp time.Time) Id {
  id := make([]byte, 12)
  copy(id[0:], timestampBytes(timestamp))
  copy(id[4:], pidBytes())
  copy(id[6:], machineIdBytes(3))
  copy(id[9:], randomBytes(2))
  id[11] = checksumByte(id[:11])
  return Id(id)
}

func (self Id) IsNil() bool { return (self == nil || len(self) == 0) }
func (self Id) ChecksumValid() bool { return checksumValid(self[:11], self[11]) }
func (self Id) IsValid() bool { return !self.IsNil() && self.ChecksumValid() } 

// Id Chain Methods 
func (self Id) Prefix(prefix string) Id {
  return self
}

