package markdown

import (
	"sort"
	"strings"
)

type SortDirection int

const (
	noDirection      = SortDirection(0)
	ASCENDING_ORDER  = SortDirection(1)
	DESCENDING_ORDER = SortDirection(-1)
)

func (s SortDirection) StringCompare(a, b string) int {
	return int(s) * strings.Compare(a, b)
}

type CompareColumnValuesFn func(s1, s2 string) int

type sortedMatrix struct {
	data [][]string
	fns  []CompareColumnValuesFn
}

func (s *sortedMatrix) Len() int {
	return len(s.data)
}

func (s *sortedMatrix) Less(i, j int) bool {
	// assumes 0 <= len(s.fns) <= len(s.data)
	for c, fn := range s.fns {
		comparison := fn(s.data[i][c], s.data[j][c])
		if comparison == 0 {
			continue
		}
		return comparison < 0
	}
	return false
}

func (s *sortedMatrix) Swap(i, j int) {
	// assumes for all i that s.data[i] is the same
	for c := 0; c < len(s.data[0]); c++ {
		s.data[i][c], s.data[j][c] = s.data[j][c], s.data[i][c]
	}
}

func SortTable(data [][]string, fns ...CompareColumnValuesFn) {
	if len(fns) == 0 {
		return
	}
	sort.Stable(&sortedMatrix{
		data: data,
		fns:  fns,
	})
}
