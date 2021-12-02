# venv

[![GoDoc](https://godoc.org/github.com/georgettica/venv?status.svg)](https://godoc.org/github.com/georgettica/venv)

This is a Go library to abstract access to environment variables.  
Like [spf13/afero][afero] or [blang/vfs][vfs], but for the env.

## Usage

```go
package main

import (
	"fmt"
	"github.com/georgettica/venv"
)

func main() {
	var e venv.Env

	// Use the real environment

	e = venv.OS()
	fmt.Printf("Hello, %s!\n", e.Getenv("USER"))

	// Or use a mock

	e = venv.Mock()
	e.Setenv("USER", "fred")
	fmt.Printf("Hello, %s!\n", e.Getenv("USER"))
}
```

## License

MIT.
