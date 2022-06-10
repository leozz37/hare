# Hare Sockets 🐇

<img align="right" width="159px" src="./resources/images/small-icon.png">

![Go](https://github.com/leozz37/hare/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/leozz37/hare/branch/main/graph/badge.svg?token=QC44PEpHRi)](https://codecov.io/gh/leozz37/hare)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/leozz37/hare)](https://goreportcard.com/report/github.com/leozz37/hare)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/4a5488c02fd6481d826512536181a1eb)](https://app.codacy.com/gh/leozz37/hare?utm_source=github.com&utm_medium=referral&utm_content=leozz37/hare&utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/97a96b7d488b201aab7c/maintainability)](https://codeclimate.com/github/leozz37/hare/maintainability)
[![GoDoc](https://pkg.go.dev/badge/github.com/leozz37/hare?status.svg)](https://pkg.go.dev/github.com/leozz37/hare?tab=doc)
[![Join the chat at https://gitter.im/hare-sockets/community](https://badges.gitter.im/hare-sockets/community.svg)](https://gitter.im/hare-sockets/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Release](https://img.shields.io/github/v/release/leozz37/hare)](https://github.com/leozz37/hare/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Hare is an user-friendly package for sockets in Golang and a CLI tool for sockets interaction. You can send and listen to TCP connections with a few lines of code or commands.

## Contents

- [Installation](#installation)
- [Quickstart](#quickstart)
- [Documentation](#documentation)
- [Examples](#examples)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## 🖥️ Installation

Installation guide for the CLI Tool and Golang Library.

### 💻 CLI Tool

To install the CLI tool, you can install it through [Homebrew](https://brew.sh/):

```shell
$ brew tap leozz37/hare

$ brew install hare
```

Or you can install manually with the `Makefile` script:

```shell
$ make install
```

### 🐇 Golang Lib

First, you need [Go](https://golang.org/) (version 1.12+ is required), then you can install Hare:

```shell
$ go get -u "github.com/leozz37/hare"
```

Import it in your code:

```shell
import "github.com/leozz37/hare"
```

## 🍕 Quickstart

Quick start for the CLI Tool and the Golang Library.

### 💻 CLI Tool

To use the CLI tool, these are the flags:

```
  -d string
        Data to be sended
  -h string
        Host address to bo operated (default "localhost")
  -l    Listen to a given address
  -p string
        Port address to bo operated         [REQUIRED]
  -s    Send message to a given address
```

You can run the `--help`  flag:

```shell
$ hare --help
```

To `Listen` to port `3000` for example, run:

```shell
$ hare -l -p 3000
```

To `Send` a payload with the message `Hello World` to port `3000` for example, run:

```shell
$ hare -s -p 3000 -d 'Hello World'
```

![cli-example](resources/images/cli-example.png)

### 🐇 Golang Lib

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

[Sample code](./examples/listen.go) for listening a port:

```go
package main

import (
    "fmt"

    "github.com/leozz37/hare"
)

func main() {
    r, _ := hare.Listen("3000")

    for {
        if r.HasNewMessages() {
            fmt.Println(r.GetMessage())
        }
    }
}
```

## 📖 Documentation

The library consists of two features: **listen** and **send** to a given port. You can check the full documentation on [Godoc](https://pkg.go.dev/github.com/leozz37/hare#section-documentation).

### Send

Receives a `port` and a `message`, both as `string` and returns an `error` (if something goes wrong).

```go
func Send(port, message string) error;
```

Usage example:

```go
func main() {
    err := hare.Send(3000, "Hello, World")
    if err != nil {
        panic(err)
    }
}
```

---

### Listen

Receives a `port` as `string` and returns a `Listener` struct and an `error` (if something goes wrong).

```go
func Listen(port string) (*Listener, error);
```

Usage example:

```go
func main() {
    r, _ := hare.Listen("3000")
    l, _ := hare.listen("3001")

    for {
        if r.HasNewMessages() {
            fmt.Println(r.GetMessage())
        } else if l.HasNewMessages() {
            fmt.Println(l.GetMessage())
        }
    }
```

---

### Listener

The **Listener** struct returned by `Listen()` function has the following fields:

```go
type Listener struct {
    SocketListener net.Listener
    HasNewMessages func() bool
    GetMessage     func() string
    Stop           func()
}
```

`SocketListener` is the socket connection.

```go
listener.SocketListener, _ = net.Listen("tcp", "localhost:" + port)
```

`HasNewMessages()` function returns a `bool` being `true` with there's a new message:

```go
func main() {
    r, _ := hare.Listen("3000")

    if r.HasNewMessages() {
        fmt.Println("There's a new message!")
    }
}
```

`GetMessage()` function returns a `string` with the last message received on the socket:

```go
func main() {
    r, _ := hare.Listen("3000")

    if r.HasNewMessages() {
        fmt.Println(r.GetMessage())
    }
}
```

`Stop()` function closes the listener connection:

```go
func main() {
    r, _ := hare.Listen("3000")
    hare.Send("3000", "Hey beauty")

    r.Stop()

    err := Send("3000", "This should fails")
    if err != nil {
        panic(err)
    }
}
```

## 📙 Examples

You can check the [example](./examples) for code usages, like [send](./examples/send.go) and [listen](./examples/listen.go) samples.

Since Hare only listens and send messages, here's a complete example:

```go
package main

import (
    "fmt"
    "time"

    "github.com/leozz37/hare"
)

func listenSockets(port string) {
    r, _ := hare.Listen(port)

    for {
        if r.HasNewMessages() {
            fmt.Println(r.GetMessage())
        }
    }
}

func main() {
    go listenSockets("3000")
    go listenSockets("3001")

    for {
        hare.Send("3000", "Hello port 3000")
        hare.Send("3001", "Hello port 3001")
        time.Sleep(time.Second)
    }
}
```

## 🧪 Testing

To run the test suite, you can run with:

```shell
$ go test
```

If you want a more detailed report with coverage and an `coverage.out` file, do the following:

```shell
$ go test -v -covermode=count -coverprofile=coverage.out
```

## 🤝 Contributing

A full guideline about contributing to Hare can be found in the [CONTRIBUTING.md](./CONTRIBUTING.md) file.

## ⚖️ License

Hare is released under the [MIT License](./LICENSE).
