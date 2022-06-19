<div align="center">
<h1>Assert</h1>

[Assert](https://github.com/goodjobtech/assert), is a testing library built on top of Generics


</div>

## Prerequisites

* Go: +1.18

## Installation

With Go module support (Go 1.11+), simply add the following import

```go
import "github.com/goodjobtech/assert"
```

Otherwise, to install *assert*, run the following command:

```shell
go get -u github.com/goodjobtech/assert
```

## Usage

```go
package test

import (
	"github.com/goodjobtech/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	assert.Equal(t, 1, 1)

	assert.NotEqual(t, "abc", "def")

	assert.Nil(t, nil)

	assert.NotNil(t, "")

	assert.Contains(t, []int{1, 2, 3}, 1)

	assert.NotContains(t, []int{1, 2, 3}, 4)

	assert.True(t, true)

	assert.False(t, false)
}
```


## Contributing

All kinds of pull request and feature requests are welcomed!

## License

Asserts's source code is licensed under [MIT License](https://choosealicense.com/licenses/mit/).
