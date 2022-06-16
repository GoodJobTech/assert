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

Otherwise, to install the assert, run the following command:

```shell
$ go get -u github.com/goodjobtech/assert
```

## Usage

```go
package something

import "testing"

func TestSomething(t *testing.T) {
	Equal(t, 1, 1)

	NotEqual(t, "abc", "def")

	Nil(t, nil)

	NotNil(t, "")

	Contains(t, []int{1, 2, 3}, 1)

	NotContains(t, []int{1, 2, 3}, 4)

	True(t, true)

	False(t, false)
}

```


## Contributing

All kinds of pull request and feature requests are welcomed!

## License

Asserts's source code is licensed under [MIT License](https://choosealicense.com/licenses/mit/).
