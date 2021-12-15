package markdown_test

import (
	. "github.com/fbiville/markdown-table-formatter/pkg/markdown"
	"reflect"
	"testing"
)

func TestSorting(st *testing.T) {

	st.Run("does not sort by default", func(t *testing.T) {
		data := [][]string{
			{"b"},
			{"a"},
			{"c"},
		}

		SortTable(data)

		expected := [][]string{
			{"b"},
			{"a"},
			{"c"},
		}
		if !reflect.DeepEqual(expected, data) {
			t.Errorf("Expected %v, got %v", expected, data)
		}
	})

	st.Run("does not sort by default, with nil slices of sorters", func(t *testing.T) {
		var sortFns []SortFunction
		data := [][]string{
			{"b"},
			{"a"},
			{"c"},
		}

		SortTable(data, sortFns...)

		expected := [][]string{
			{"b"},
			{"a"},
			{"c"},
		}
		if !reflect.DeepEqual(expected, data) {
			t.Errorf("Expected %v, got %v", expected, data)
		}
	})

	st.Run("sorts single column with function", func(t *testing.T) {
		data := [][]string{
			{"b"},
			{"a"},
			{"c"},
		}

		SortTable(data, ASCENDING_ORDER.StringCompare(0))

		expected := [][]string{
			{"a"},
			{"b"},
			{"c"},
		}
		if !reflect.DeepEqual(expected, data) {
			t.Errorf("Expected %v, got %v", expected, data)
		}
	})

	st.Run("sorts columns with multiple functions", func(t *testing.T) {
		data := [][]string{
			{"b", "b1"},
			{"a", "a1"},
			{"a", "z1"},
		}

		SortTable(data, ASCENDING_ORDER.StringCompare(0), DESCENDING_ORDER.StringCompare(1))

		expected := [][]string{
			{"a", "z1"},
			{"a", "a1"},
			{"b", "b1"},
		}
		if !reflect.DeepEqual(expected, data) {
			t.Errorf("Expected %v, got %v", expected, data)
		}
	})

	st.Run("sorts columns with multiple functions in different column order", func(t *testing.T) {
		data := [][]string{
			{"b", "z1"},
			{"a", "b1"},
			{"z", "b1"},
		}

		SortTable(data, ASCENDING_ORDER.StringCompare(1), DESCENDING_ORDER.StringCompare(0))

		expected := [][]string{
			{"z", "b1"},
			{"a", "b1"},
			{"b", "z1"},
		}
		if !reflect.DeepEqual(expected, data) {
			t.Errorf("Expected %v, got %v", expected, data)
		}
	})

	st.Run("sorts according to first column, re-arranges the rest", func(t *testing.T) {
		data := [][]string{
			{"b", "b1"},
			{"a", "a1"},
			{"c", "c1"},
		}

		SortTable(data, ASCENDING_ORDER.StringCompare(0))

		expected := [][]string{
			{"a", "a1"},
			{"b", "b1"},
			{"c", "c1"},
		}
		if !reflect.DeepEqual(expected, data) {
			t.Errorf("Expected %v, got %v", expected, data)
		}
	})

	st.Run("sorts according to second column, re-arranges the rest", func(t *testing.T) {
		data := [][]string{
			{"b", "b1"},
			{"a", "a1"},
			{"c", "c1"},
		}

		SortTable(data, DESCENDING_ORDER.StringCompare(1))

		expected := [][]string{
			{"c", "c1"},
			{"b", "b1"},
			{"a", "a1"},
		}
		if !reflect.DeepEqual(expected, data) {
			t.Errorf("Expected %v, got %v", expected, data)
		}
	})

	st.Run("sorts according to first n columns, re-arranges the rest", func(t *testing.T) {
		data := [][]string{
			{"b", "b1", "b2"},
			{"a", "a1", "a2"},
			{"a", "z1", "z2"},
		}

		SortTable(data, ASCENDING_ORDER.StringCompare(0), DESCENDING_ORDER.StringCompare(1))

		expected := [][]string{
			{"a", "z1", "z2"},
			{"a", "a1", "a2"},
			{"b", "b1", "b2"},
		}
		if !reflect.DeepEqual(expected, data) {
			t.Errorf("Expected %v, got %v", expected, data)
		}
	})

	st.Run("sorts according to some unordered columns, re-arranges the rest", func(t *testing.T) {
		data := [][]string{
			{"b", "b1", "b2"},
			{"a", "a1", "b2"},
			{"a", "z1", "z2"},
		}

		SortTable(data, DESCENDING_ORDER.StringCompare(2), ASCENDING_ORDER.StringCompare(0))

		expected := [][]string{
			{"a", "z1", "z2"},
			{"a", "a1", "b2"},
			{"b", "b1", "b2"},
		}
		if !reflect.DeepEqual(expected, data) {
			t.Errorf("Expected %v, got %v", expected, data)
		}
	})
}
