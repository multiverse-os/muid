package id

import (
	"sort"
)

type sorter []Id

func (self sorter) Len() int {
	return len(self)
}

func (self sorter) Less(i, j int) bool {
	return self[i].Compare(self[j]) < 0
}

func (self sorter) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func Sort(ids []Id) {
	sort.Sort(sorter(ids))
}
