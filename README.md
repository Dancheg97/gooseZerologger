# Goose to Zerolog

Adapter from plessy goose to zerolog. Allows to transform zerolog `Logger` to
`goose` compatible logging tool, by using tiny adapter.

# Installation

Go get:

```cmd
go get github.com/Dancheg97/gooseZerologger
```

# Usage

Example usage:

```go
package main

import (
	"github.com/Dancheg97/gooseZerologger/adapter"
	"github.com/pressly/goose"
	"github.com/rs/zerolog"
)

func main() {
	adapter := adapter.Get(&zerolog.Logger{})
	goose.SetLogger(adapter)
}

```
