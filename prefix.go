package muid

func (self Id) Prefix(prefix string) Id {
  newId := make([]byte, len(self)+len(prefix))
  copy(newId[:len(prefix)], prefixBytes(prefix))
  copy(newId[len(prefix):], self)
  return Id(newId)
}

func prefixBytes(prefix string) []byte { return []byte(prefix) }
