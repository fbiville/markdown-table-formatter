# Markdown table formatter

## Import

```sh
 $ go get github.com/fbiville/markdown-table-formatter/pkg/markdown 
```

## Usage

### Basic

```go
package some_package

import (
	"fmt"
	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
)

func someFunction() {
	basicTable, err := markdown.NewTableFormatterBuilder().
		Build("column 1", "column 2", "column 3").
		Format([][]string{
			{"so", "much", "fun"},
			{"don't", "you", "agree"},
		})
	
	if err != nil {
		// ... do your thing
	}
	fmt.Print(basicTable)
}
```

The output will be:
```markdown
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| so | much | fun |
| don't | you | agree |
```

### Pretty-print

```go
package some_package

import (
	"fmt"
	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
)

func someFunction() {
	prettyPrintedTable, err := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build("column 1", "column 2", "column 3").
		Format([][]string{
			{"so", "much", "fun"},
			{"don't", "you", "agree"},
		})
	
	if err != nil {
		// ... do your thing
	}
	fmt.Print(prettyPrintedTable)
}
```

The output will be:
```markdown
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| so       | much     | fun      |
| don't    | you      | agree    |
```

### Sort

Sorting by columns is disabled by default.
It is available to all formatters.

#### Alphabetical order

When using `markdown.ASCENDING_ORDER` or `markdown.DESCENDING_ORDER`, the associated string comparison algorithm
is applied to **every** column, in column order.

##### Ascending order

```go
table, err := markdown.NewTableFormatterBuilder().
    WithAlphabeticalSortIn(markdown.ASCENDING_ORDER).
    Build("column 1", "column 2", "column 3").
    Format([][]string{
        {"don't", "you", "know"},
        {"so", "much", "fun"},
        {"don't", "you", "agree"},
    })

if err != nil { 
    // ... do your thing 
}
fmt.Print(table)
```

The output will be:
```markdown
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| don't | you | agree |
| don't | you | know |
| so | much | fun |
```

##### Descending order

```go
table, err := markdown.NewTableFormatterBuilder().
    WithAlphabeticalSortIn(markdown.DESCENDING_ORDER).
    // ... same as above ...
    
fmt.Print(basicTable)
```

The output will be:
```markdown
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| so | much | fun |
| don't | you | know |
| don't | you | agree |
```

#### Custom sort functions

There can be at most `N` custom sort functions (where `N` is the number of columns).

If the array is made of 3 columns e.g., 

```go
WithAlphabeticalSortIn(markdown.ASCENDING_ORDER)
```

is effectively equivalent to:

```go
WithCustomSort(markdown.ASCENDING_ORDER.StringCompare(0),
               markdown.ASCENDING_ORDER.StringCompare(1),
               markdown.ASCENDING_ORDER.StringCompare(2))
```

This is **not** the same as:

```go
WithCustomSort(markdown.ASCENDING_ORDER.StringCompare(0))
```

Indeed, providing a single function triggers a comparison between values of the first column, the other columns' values
are not compared.

In other words, the following program:

```go
table, err := markdown.NewTableFormatterBuilder().
    WithCustomSort(markdown.ASCENDING_ORDER.StringCompare(0)).
    // equivalent to:
    // WithCustomSort(
    //  SortFunction{
    //      Fn: func(a, b string) int {
    //          // you can plug *any* logic here
    //          return strings.Compare(a, b)
    //      },
    //      Column: 0,
    // })
    WithPrettyPrint().
    Build("column 1", "column 2", "column 3").
    Format([][]string{
        {"don't", "you", "know"},
        {"so", "much", "fun"},
        {"don't", "you", "agree"},
    })

if err != nil { 
    // ... do your thing 
}

fmt.Print(table)
```

... will yield:
```markdown
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| don't    | you      | know     |
| don't    | you      | agree    |
| so       | much     | fun      |
```

The order of custom sort function is the order of **execution**.
You are free to sort columns by any order:

```go
table, err := markdown.NewTableFormatterBuilder().
    // sort third column first, then first column
    WithCustomSort(markdown.ASCENDING_ORDER.StringCompare(2), markdown.DESCENDING_ORDER.StringCompare(0)).
    WithPrettyPrint().
    Build("column 1", "column 2", "column 3").
    Format([][]string{
        {"don't", "you", "know"},
        {"so", "much", "agree"},
        {"don't", "you", "agree"},
    })

if err != nil { 
    // ... do your thing 
}

fmt.Print(table)
```

... will yield:

```markdown
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| so       | much     | agree    |
| don't    | you      | agree    |
| don't    | you      | know     |
```