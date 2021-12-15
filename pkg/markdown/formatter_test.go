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
