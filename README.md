# hare 🐇

User friendly socket lib for Golang

## Installation

Install Hare:

```shell
$ go get "github.com/leozz37/hare
```

Import it in your code:

```shell
import "github.com/leozz37/hare"
```

## Quick start

You can check the [example code]() for listening:

```go
package main

import (
    "fmt"
    "github.com/leozz37/hare"
)

func main() {
    r, err := hare.Listen(3000)
    if err {
        panic(err)
    }

    if r.HasMessage() {
        fmt.Println(r.LastMessage())
    }
}
```

Example code for sending payloads:

```go
package main

import (
    "github.com/leozz37/hare"
)

func main() {
    hare.Send(3000, "Hello, World")
}
```
