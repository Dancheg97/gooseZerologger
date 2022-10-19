# Goose to Zerolog

Adapter from plessy goose to zerolog.

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
