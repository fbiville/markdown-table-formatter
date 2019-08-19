package markdown_test

import (
	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
	. "github.com/onsi/gomega"
	"strings"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Default Markdown table markdown", func() {

	var formatter markdown.TableFormatter

	BeforeEach(func() {
		formatter = markdown.NewTableFormatterBuilder().
			Build("column 1", "column 2", "column 3")
	})

	It("formats accordingly", func() {
		output, err := formatter.Format([][]string{
			{"value 1", "val 2", "v 3"},
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(output).To(Equal(strings.TrimLeft(`
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| value 1 | val 2 | v 3 |
`, "\n")))
	})

	It("fails to format if the data dimension does not match the headers", func() {
		_, err := formatter.Format([][]string{
			{"value 1", "value 2"},
		})

		Expect(err).To(MatchError("expected 3 column(s), row number 1 got 2"))
	})
})

var _ = Describe("Pretty Markdown table markdown", func() {

	var formatter markdown.TableFormatter

	BeforeEach(func() {
		formatter = markdown.NewTableFormatterBuilder().
			WithPrettyPrint().
			Build("column 1", "column 2", "column 3")
	})

	It("pretty-prints accordingly", func() {
		output, err := formatter.Format([][]string{
			{"value 1", "val 2", "v 3"},
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(output).To(Equal(strings.TrimLeft(`
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| value 1  | val 2    | v 3      |
`, "\n")))
	})

	It("fails to format if the data dimension does not match the headers", func() {
		_, err := formatter.Format([][]string{
			{"value 1", "value 2"},
		})

		Expect(err).To(MatchError("expected 3 column(s), row number 1 got 2"))
	})
})
