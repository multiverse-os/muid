package muid

import (
  "bytes"
	"sort"
)

// TODO: Ideally the resulting strings should be sortable and this may not be
// necessary and it could probably be written better but this is on the back
// burner for now
type sorter []Id

func (self sorter) Len() int { return len(self) }
func (self sorter) Less(i, j int) bool { return self[i].Compare(self[j]) < 0 }
func (self sorter) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

func Sort(ids []Id) { sort.Sort(sorter(ids)) }

// TODO: Would be best to probably just compare the time portion of the ID,
// likely less resources, and faster which is important for sorting. 
func (self Id) Compare(other Id) int {
	return bytes.Compare(self[:], other[:])
}
