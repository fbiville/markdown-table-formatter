package markdown

import (
	"fmt"
	"strings"
)

type TableFormatter interface {
	Format(data [][]string) (string, error)
}

func newDefaultTableFormatter(headers []string) TableFormatter {
	return &defaultTableFormatter{
		config: &config{headers: headers},
	}
}

func newPrettyTableFormatter(headers []string) TableFormatter {
	return &prettyTableFormatter{
		config: &config{headers: headers},
	}
}

type defaultTableFormatter struct {
	config *config
}

func (formatter *defaultTableFormatter) Format(data [][]string) (string, error) {
	builder := &strings.Builder{}
	appendHeaders(builder, formatter.config.headers)
	for rowIndex, row := range data {
		if err := formatter.config.validateRow(rowIndex, row); err != nil {
			return "", err
		}
		builder.WriteString(joinValues(row))
	}
	return builder.String(), nil
}

type prettyTableFormatter struct {
	config *config
}

type prettyTable struct {
	widths  []int
	content [][]string
}

func (formatter *prettyTableFormatter) Format(data [][]string) (string, error) {
	prettyTable, err := formatter.preComputeFormattedData(data)
	if err != nil {
		return "", err
	}
	builder := &strings.Builder{}
	appendHeaders(builder, formatter.config.headers)
	for _, row := range prettyTable.content {
		builder.WriteString(joinValues(replacePadded(row, prettyTable.widths)))
	}
	return builder.String(), nil
}

func (formatter *prettyTableFormatter) preComputeFormattedData(data [][]string) (*prettyTable, error) {
	var widths []int
	for _, header := range formatter.config.headers {
		widths = append(widths, len(header))
	}
	for rowIndex, row := range data {
		if err := formatter.config.validateRow(rowIndex, row); err != nil {
			return nil, err
		}
		for columnIndex, cell := range row {
			cellLength := len(cell)
			if cellLength > widths[columnIndex] {
				widths[columnIndex] = cellLength
			}
		}
	}
	return &prettyTable{widths: widths, content: data}, nil
}

func appendHeaders(builder *strings.Builder, headers []string) {
	builder.WriteString(joinValues(headers))
	builder.WriteString(joinValues(replaceRepeated(headers, "-")))
}

type config struct {
	headers []string
}

func (config *config) validateRow(rowIndex int, row []string) error {
	headerCount := len(config.headers)
	rowLength := len(row)
	if rowLength != headerCount {
		return fmt.Errorf("expected %d column(s), row number %d got %d", headerCount, 1+rowIndex, rowLength)
	}
	return nil
}

func joinValues(values []string) string {
	return fmt.Sprintf("| %s |\n", strings.Join(values, " | "))
}

func replaceRepeated(items []string, replacementSymbol string) []string {
	var result []string
	for _, item := range items {
		result = append(result, strings.Repeat(replacementSymbol, len(item)))
	}
	return result
}

func replacePadded(items []string, widths []int) []string {
	var result []string
	for columnIndex, item := range items {
		format := fmt.Sprintf("%%-%ds", widths[columnIndex])
		result = append(result, fmt.Sprintf(format, item))
	}
	return result
}
