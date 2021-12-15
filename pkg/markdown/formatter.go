package markdown

import (
	"fmt"
	"strings"
)

type TableFormatter interface {
	Format(data [][]string) (string, error)
}

func newDefaultTableFormatter(headers []string, sortFns []SortFunction) TableFormatter {
	return &defaultTableFormatter{
		config:  &config{headers: headers},
		sortFns: sortFns,
	}
}

func newPrettyTableFormatter(headers []string, sortFns []SortFunction) TableFormatter {
	return &prettyTableFormatter{
		config:  &config{headers: headers},
		sortFns: sortFns,
	}
}

type defaultTableFormatter struct {
	config  *config
	sortFns []SortFunction
}

func (dtf *defaultTableFormatter) Format(data [][]string) (string, error) {
	// this could be checked before but would require incompatible API change
	if err := checkSortingConfiguration(dtf.sortFns, dtf.config.headers); err != nil {
		return "", err
	}

	SortTable(data, dtf.sortFns...)
	builder := &strings.Builder{}
	appendHeaders(builder, dtf.config.headers)
	for rowIndex, row := range data {
		if err := dtf.config.validateRow(rowIndex, row); err != nil {
			return "", err
		}
		builder.WriteString(joinValues(row))
	}
	return builder.String(), nil
}

type prettyTableFormatter struct {
	config  *config
	sortFns []SortFunction
}

type prettyTable struct {
	widths  []int
	content [][]string
}

func (ptf *prettyTableFormatter) Format(data [][]string) (string, error) {
	// this could be checked before but would require incompatible API change
	if err := checkSortingConfiguration(ptf.sortFns, ptf.config.headers); err != nil {
		return "", err
	}
	SortTable(data, ptf.sortFns...)
	prettyTable, err := ptf.preComputeFormattedData(data)
	if err != nil {
		return "", err
	}
	widths := prettyTable.widths
	builder := &strings.Builder{}
	appendHeaders(builder, replacePadded(ptf.config.headers, widths))
	for _, row := range prettyTable.content {
		builder.WriteString(joinValues(replacePadded(row, widths)))
	}
	return builder.String(), nil
}

func (ptf *prettyTableFormatter) preComputeFormattedData(data [][]string) (*prettyTable, error) {
	var widths []int
	for _, header := range ptf.config.headers {
		widths = append(widths, len(header))
	}
	for rowIndex, row := range data {
		if err := ptf.config.validateRow(rowIndex, row); err != nil {
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

func checkSortingConfiguration(sortFns []SortFunction, headers []string) error {
	if len(sortFns) > len(headers) {
		return fmt.Errorf("expected at most %d sort functions, %d given", len(headers), len(sortFns))
	}
	columnCount := make(map[int]struct{}, len(headers))
	for _, sortFn := range sortFns {
		if sortFn.Column >= len(headers) {
			return fmt.Errorf("expected column index to be between 0 included and %d excluded, got %d", len(headers), sortFn.Column)
		}
		if _, found := columnCount[sortFn.Column]; found {
			return fmt.Errorf("expected at most 1 sort function for column index %d, found at least 2", sortFn.Column)
		}
		columnCount[sortFn.Column] = struct{}{}
	}

	return nil
}
