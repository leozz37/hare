# hare 🐇

<img align="right" width="159px" src="./resources/images/small-icon.png">

![Go](https://github.com/leozz37/hare/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/leozz37/hare)](https://goreportcard.com/report/github.com/leozz37/hare)
[![GoDoc](https://pkg.go.dev/badge/github.com/leozz37/hare?status.svg)](https://pkg.go.dev/github.com/leozz37/hare?tab=doc)

User friendly socket lib for Golang

## Contents

- [Hare](#hare-)
  - [Contents](#contents)
  - [Installation](#installation)
  - [Quick start](#quick-start)
  - [Documentation](#documentation)
  - [Examples](#examples)
  - [Testing](#testing)

## Installation

Install Hare:

```shell
$ go get "github.com/leozz37/hare"
```

Import it in your code:

```shell
import "github.com/leozz37/hare"
```

## Quick start

[Example code](./examples/send.go) for sending payloads:

```go
package main

import (
    "github.com/leozz37/hare"
)

func main() {
    hare.Send(3000, "Hello, World")
}
```

[Example code](./examples/listen.go) for listening port:

```go
package main

import (
    "fmt"

    "github.com/leozz37/hare"
)

func main() {
    r := hare.Listen("3000")

    for {
        if r.HasNewMessages {
            fmt.Println(hare.GetMessage())
        }
    }
}
```

## Documentation

TODO: Implement documentation

## Examples

TODO: Implement examples

## Testing

TODO: Implement testing instructions
