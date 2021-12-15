package markdown

func NewTableFormatterBuilder() *tableFormatterBuilder {
	return &tableFormatterBuilder{}
}

type tableFormatterBuilder struct {
	alphaSortDirection SortDirection
	customSortFns      []SortFunction
}

func (tfb *tableFormatterBuilder) WithAlphabeticalSortIn(direction SortDirection) *tableFormatterBuilder {
	tfb.alphaSortDirection = direction
	return tfb
}

func (tfb *tableFormatterBuilder) WithCustomSort(sortFns ...SortFunction) *tableFormatterBuilder {
	tfb.customSortFns = sortFns
	return tfb
}

func (tfb *tableFormatterBuilder) Build(headers ...string) TableFormatter {
	return newDefaultTableFormatter(
		headers,
		sortFunctions(tfb.customSortFns, tfb.alphaSortDirection, headers),
	)
}

func (tfb *tableFormatterBuilder) WithPrettyPrint() *prettyTableFormatterBuilder {
	return &prettyTableFormatterBuilder{
		alphaSortDirection: tfb.alphaSortDirection,
		customSortFns:      tfb.customSortFns,
	}
}

type prettyTableFormatterBuilder struct {
	alphaSortDirection SortDirection
	customSortFns      []SortFunction
}

func (ptfb *prettyTableFormatterBuilder) Build(headers ...string) TableFormatter {
	return newPrettyTableFormatter(
		headers,
		sortFunctions(ptfb.customSortFns, ptfb.alphaSortDirection, headers),
	)
}

func sortFunctions(customSortFns []SortFunction, alphaSortDirection SortDirection, headers []string) []SortFunction {
	if len(customSortFns) > 0 {
		return customSortFns
	}
	if alphaSortDirection != noDirection {
		result := make([]SortFunction, len(headers))
		for i := 0; i < len(headers); i++ {
			result[i] = alphaSortDirection.StringCompare(i)
		}
		return result
	}
	return nil
}
