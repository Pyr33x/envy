# Envy
Typed and fallback-safe environment variables access in Go

[![Workflows](https://img.shields.io/github/actions/workflow/status/pyr33x/envy/test.yml?label=tests&labelColor=blue&color=gray)](https://github.com/Pyr33x/envy/actions)
[![Git Tag](https://img.shields.io/github/v/tag/pyr33x/envy?include_prereleases&sort=date&labelColor=blue&color=gray)](https://github.com/Pyr33x/envy/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/pyr33x/envy.svg)](https://pkg.go.dev/github.com/pyr33x/envy)
[![Coverage](https://codecov.io/gh/pyr33x/envy/branch/master/graph/badge.svg)](https://codecov.io/gh/pyr33x/envy)
[![License](https://img.shields.io/github/license/pyr33x/envy?color=blue)](https://github.com/pyr33x/envy/blob/master/LICENSE)

## Install
```sh
go get github.com/pyr33x/envy
```

## Usage
Make sure you've [loaded](https://github.com/joho/godotenv) your environment variables.

### With Fallback

```go
package main

import (
	"fmt"

	"github.com/pyr33x/envy"
)

func main() {
	b := envy.GetBool("A", true)
	fmt.Println(b) // returns value from env | if not set *true* (fallback)

	f := envy.GetFloat64("B", 10.5)
	fmt.Println(f) // returns value from env | if not set *10.5* (fallback)

	u := envy.GetUint64("C", 10)
	fmt.Println(u) // returns value from env | if not set *10* (fallback)

	s := envy.GetString("D", "yo")
	fmt.Println(s) // returns value from env | if not set *yo* (fallback)
}
```

### Force
```go
package main

import (
	"fmt"

	"github.com/pyr33x/envy"
)

func main() {
	s := envy.MustGetString("A")
	fmt.Println(s) // returns env value in string (exists)

	s = envy.MustGetString("T")
	fmt.Println(s) // panics, because key has not been set
}
```
