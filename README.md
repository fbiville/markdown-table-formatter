# Markdown table formatter in Golang

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
```
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
	fmt.Println(prettyPrintedTable)
}
```

The output will be:
```
| column 1 | column 2 | column 3 |
| -------- | -------- | -------- |
| so       | much     | fun      |
| don't    | you      | agree    |
```