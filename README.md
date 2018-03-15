Shellwords
===========

A golang library for splitting command-line strings into tokens like a POSIX or Windows shell would.

## Installation

```bash
go get -u github.com/buildkite/shellwords
```

## Usage

```go
package main

import (
  "github.com/buildkite/shellwords"
  "fmt"
)

func main() {
  words := shellwords.SplitPosix(`/usr/bin/bash -e -c "llamas are the \"best\" && echo 'alpacas'"`)
  for _, word := range words {
    fmt.Println(word)
  }
}
```

## License

Licensed under MIT license, in `LICENSE`.
