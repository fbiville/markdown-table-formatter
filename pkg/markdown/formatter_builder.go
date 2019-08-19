package markdown

func NewTableFormatterBuilder() *tableFormatterBuilder {
	return &tableFormatterBuilder{}
}

type tableFormatterBuilder struct {
}

func (*tableFormatterBuilder) Build(headers ...string) TableFormatter {
	return newDefaultTableFormatter(headers)
}

func (*tableFormatterBuilder) WithPrettyPrint() *prettyTableFormatterBuilder {
	return &prettyTableFormatterBuilder{}
}

type prettyTableFormatterBuilder struct {
}

func (*prettyTableFormatterBuilder) Build(headers ...string) TableFormatter {
	return newPrettyTableFormatter(headers)
}

