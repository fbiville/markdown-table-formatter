package markdown

import (
	"fmt"
	"strings"
)

type TableFormatter interface {
	Format(data [][]string) (string, error)
}

type config struct {
	headers []string
}

func NewDefaultTableFormatter(headers []string) TableFormatter {
	return &defaultTableFormatter{
		config: &config{headers: headers},
	}
}

type defaultTableFormatter struct {
	config *config
}

func (formatter *defaultTableFormatter) Format(data [][]string) (string, error) {
	builder := &strings.Builder{}
	formatter.appendHeaders(builder)
	if err := formatter.appendData(builder, data); err != nil {
		return "", err
	}
	return builder.String(), nil
}

func (formatter *defaultTableFormatter) appendHeaders(builder *strings.Builder) {
	columnNames := formatter.config.headers
	builder.WriteString(joinValues(columnNames))
	builder.WriteString(joinValues(replaceWith(columnNames, "-")))
}

func (formatter *defaultTableFormatter) appendData(builder *strings.Builder, data [][]string) error {
	headerCount := len(formatter.config.headers)
	for rowIndex, row := range data {
		rowLength := len(row)
		if rowLength != headerCount {
			return fmt.Errorf("expected %d column(s), row number %d got %d", headerCount, 1+rowIndex, rowLength)
		}
		builder.WriteString(joinValues(row))
	}
	return nil
}

func joinValues(values []string) string {
	return fmt.Sprintf("| %s |\n", strings.Join(values, " | "))
}

func replaceWith(items []string, replacementSymbol string) []string {
	var result []string
	for _, item := range items {
		result = append(result, strings.Repeat(replacementSymbol, len(item)))
	}
	return result
}
