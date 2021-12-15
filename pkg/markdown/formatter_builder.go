package markdown

func NewTableFormatterBuilder() *tableFormatterBuilder {
	return &tableFormatterBuilder{}
}

type tableFormatterBuilder struct {
	alphaSortDirection SortDirection
	customSortFns      []CompareColumnValuesFn
}

func (tfb *tableFormatterBuilder) WithAlphabeticalSortIn(direction SortDirection) *tableFormatterBuilder {
	tfb.alphaSortDirection = direction
	return tfb
}

func (tfb *tableFormatterBuilder) WithCustomSort(sortFns ...CompareColumnValuesFn) *tableFormatterBuilder {
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
	customSortFns      []CompareColumnValuesFn
}

func (ptfb *prettyTableFormatterBuilder) Build(headers ...string) TableFormatter {
	return newPrettyTableFormatter(
		headers,
		sortFunctions(ptfb.customSortFns, ptfb.alphaSortDirection, headers),
	)
}

func sortFunctions(customSortFns []CompareColumnValuesFn, alphaSortDirection SortDirection, headers []string) []CompareColumnValuesFn {
	if len(customSortFns) > 0 {
		return customSortFns
	}
	if alphaSortDirection != noDirection {
		result := make([]CompareColumnValuesFn, len(headers))
		for i := 0; i < len(headers); i++ {
			result[i] = alphaSortDirection.StringCompare
		}
		return result
	}
	return nil
}
