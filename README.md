# Hare Socket 🐇

<img align="right" width="159px" src="./resources/images/small-icon.png">

![Go](https://github.com/leozz37/hare/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/leozz37/hare/branch/main/graph/badge.svg?token=QC44PEpHRi)](https://codecov.io/gh/leozz37/hare)
[![Go Report Card](https://goreportcard.com/badge/github.com/leozz37/hare)](https://goreportcard.com/report/github.com/leozz37/hare)
[![GoDoc](https://pkg.go.dev/badge/github.com/leozz37/hare?status.svg)](https://pkg.go.dev/github.com/leozz37/hare?tab=doc)
[![Release](https://img.shields.io/github/v/release/leozz37/hare)](https://github.com/leozz37/hare/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Hare is a user-friendly lib wrapper for sockets.

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

[Sample code](./examples/send.go) for sending payloads:

```go
package main

import (
    "github.com/leozz37/hare"
)

func main() {
    hare.Send(3000, "Hello, World")
}
```

[Sample code](./examples/listen.go) for listening port:

```go
package main

import (
    "fmt"

    "github.com/leozz37/hare"
)

func main() {
    r, _ := hare.Listen("3000")

    for {
        if r.HasNewMessages {
            fmt.Println(hare.GetMessage())
        }
    }
}
```

## Documentation

You can check the full documentation on [Godoc](https://pkg.go.dev/github.com/leozz37/hare#section-documentation).

The library consists on two features: listen and send to a given port.

`Send` receives a `port` and a `message`, both as `string` and returns a `error` (if it happens).

```go
func Send(port, message string) error;
```

`Listen` receives a `port` as `string` and returns a `Listener` struct and an `error` (if it happens).

```go
func Listen(port string) (*Listener, error);
```

The `Listener` struct has the following attributes:

```go
type Listener struct {
	SocketListener net.Listener
	HasNewMessages bool
	Message        string
}
```

A `net.Listener` connection, a `bool` flag for new messages and `string` with the message.

The `GetMessage()` method returns the last message:

```go
func GetMessage() string;
```

## Examples

You can check the [example](./examples) for code usages, like [send](./examples/send.go) and [listen](./examples/listen.go) samples.

## Testing

To run the test suite, you can run with:

```shell
$ go test
```

If you want a more detailed report with coverage and an `coverage.out` file, do the following:

```shell
$ go test -v -covermode=count -coverprofile=coverage.out
```

