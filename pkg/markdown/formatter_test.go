package markdown_test

import (
	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
	"strings"
	"testing"
)

func TestDefaultRendering(st *testing.T) {
	formatter := markdown.NewTableFormatterBuilder().
		Build("column 1", "column 2", "column 3")

	st.Run("formats as-is", func(t *testing.T) {
		actual, err := formatter.Format([][]string{
			{"value 1", "val 2", "v 3"},
		})

		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		expected := strings.TrimLeft(`
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| value 1 | val 2 | v 3 |
`, "\n")
		if expected != actual {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	})

	st.Run("fails to format if the data dimension does not match the headers", func(t *testing.T) {
		_, err := formatter.Format([][]string{
			{"value 1", "value 2"},
		})

		expected := "expected 3 column(s), row number 1 got 2"
		if err == nil || err.Error() != expected {
			t.Errorf("Expected error with message %s but got %v", expected, err)
		}
	})
}

func TestPrettyPrintedRendering(st *testing.T) {

	formatter := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build("column 1", "column 2", "column 3")

	st.Run("pretty-prints accordingly", func(t *testing.T) {
		actual, err := formatter.Format([][]string{
			{"value 1", "val 2", "v 3"},
		})

		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		expected := strings.TrimLeft(`
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| value 1  | val 2    | v 3      |
`, "\n")
		if expected != actual {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	})

	st.Run("pads columns if necessary", func(t *testing.T) {
		actual, err := formatter.Format([][]string{
			{"long value 1", "val 2", "very long value 3"},
		})

		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		expected := strings.TrimLeft(`
| column 1     | column 2 | column 3          |
| ------------ | -------- | ----------------- |
| long value 1 | val 2    | very long value 3 |
`, "\n")
		if expected != actual {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	})

	st.Run("fails to format if the data dimension does not match the headers", func(t *testing.T) {
		_, err := formatter.Format([][]string{
			{"value 1", "value 2"},
		})

		expected := "expected 3 column(s), row number 1 got 2"
		if err == nil || err.Error() != expected {
			t.Errorf("Expected error with message %s but got %v", expected, err)
		}
	})
}

func TestSortedRendering(st *testing.T) {
	headers := []string{"column 1", "column 2", "column 3"}
	data := [][]string{
		{"v1", "v1.1", "v1.1.1"},
		{"w2", "w2.2", "w2.2.2"},
		{"w2", "z2.2", "w1.1.1"},
		{"z3", "z3.3", "z3.3.3"},
	}
	formatterBuilder := markdown.NewTableFormatterBuilder().WithAlphabeticalSortIn(markdown.DESCENDING_ORDER)

	st.Run("prints sorted values", func(t *testing.T) {
		formatter := formatterBuilder.Build(headers...)

		actual, err := formatter.Format(data)

		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		expected := strings.TrimLeft(`
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| z3 | z3.3 | z3.3.3 |
| w2 | z2.2 | w1.1.1 |
| w2 | w2.2 | w2.2.2 |
| v1 | v1.1 | v1.1.1 |
`, "\n")
		if expected != actual {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	})

	st.Run("pretty-prints sorted values", func(t *testing.T) {
		formatter := formatterBuilder.WithPrettyPrint().Build(headers...)

		actual, err := formatter.Format(data)

		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		expected := strings.TrimLeft(`
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| z3       | z3.3     | z3.3.3   |
| w2       | z2.2     | w1.1.1   |
| w2       | w2.2     | w2.2.2   |
| v1       | v1.1     | v1.1.1   |
`, "\n")
		if expected != actual {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	})

	st.Run("fails printing with too many sort functions", func(t *testing.T) {
		_, err := markdown.NewTableFormatterBuilder().
			WithCustomSort(strings.Compare, strings.Compare, strings.Compare).
			Build("header1", "header2").
			Format([][]string{})

		expected := "expected at most 2 sort functions, 3 given"
		if err == nil || err.Error() != expected {
			t.Errorf("Expected error with message %q, but got %v", expected, err)
		}
	})

	st.Run("fails pretty-printing with too many sort functions", func(t *testing.T) {
		_, err := markdown.NewTableFormatterBuilder().
			WithCustomSort(strings.Compare, strings.Compare, strings.Compare).
			WithPrettyPrint().
			Build("header1", "header2").
			Format([][]string{})

		expected := "expected at most 2 sort functions, 3 given"
		if err == nil || err.Error() != expected {
			t.Errorf("Expected error with message %q, but got %v", expected, err)
		}
	})
}
