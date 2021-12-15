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

func (s SortDirection) StringCompare(column int) SortFunction {
	return SortFunction{
		Fn: func(a, b string) int {
			return int(s) * strings.Compare(a, b)
		},
		Column: column,
	}
}

type SortFunction struct {
	Fn     func(s1, s2 string) int
	Column int
}

type sortedMatrix struct {
	data          [][]string
	sortFunctions []SortFunction
}

func (s *sortedMatrix) Len() int {
	return len(s.data)
}

func (s *sortedMatrix) Less(i, j int) bool {
	for _, sortFunction := range s.sortFunctions {
		comparison := sortFunction.Fn(s.data[i][sortFunction.Column], s.data[j][sortFunction.Column])
		if comparison == 0 {
			continue
		}
		return comparison < 0
	}
	return false
}

func (s *sortedMatrix) Swap(i, j int) {
	// assumes for all i that len(s.data[i]) is the same
	for c := 0; c < len(s.data[0]); c++ {
		s.data[i][c], s.data[j][c] = s.data[j][c], s.data[i][c]
	}
}

func SortTable(data [][]string, fns ...SortFunction) {
	if len(fns) == 0 {
		return
	}
	sort.Stable(&sortedMatrix{
		data:          data,
		sortFunctions: fns,
	})
}
