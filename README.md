# Furagu
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mrmarble/furagu)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrmarble/furagu)](https://goreportcard.com/report/github.com/mrmarble/furagu)
![License](https://img.shields.io/badge/License-MIT-blue)
[![Go Reference](https://pkg.go.dev/badge/github.com/mrmarble/furagu.svg)](https://pkg.go.dev/github.com/mrmarble/furagu)



Furagu is a lightweight Go library that extends the standard `flag` package to support environment variable defaults for command-line flags. It is designed to make CLI applications more flexible and configurable by allowing flag values to be set via environment variables.

## Features

- Simple API for defining flags with environment variable support
- Supports `bool`, `int`, `string`, `float64`, and `time.Duration` types
- Automatic usage message generation with environment variable info
- Panics on invalid environment variable values for safety

## Installation

```bash
go get github.com/mrmarble/furagu
```

## Usage

```go
package main

import (
  "fmt"
  "github.com/mrmarble/furagu"
)

var (
  debug = furagu.Bool("debug", false, "Enable debug mode")
  port  = furagu.Int("port", 8080, "Port to listen on")
)

func main() {
  // Parse flags, using "MYAPP" as the environment variable prefix
  furagu.Parse("MYAPP")

  fmt.Println("Debug:", *debug)
  fmt.Println("Port:", *port)
}
```

You can set environment variables like `MYAPP_DEBUG` or `MYAPP_PORT` to override the default values.

## Example

```bash
export MYAPP_DEBUG=true
export MYAPP_PORT=9000
go run main.go
```

## Documentation

See the [full documentation](https://pkg.go.dev/github.com/mrmarble/furagu) for detailed usage and API reference.

## Contributing

Contributions are welcome! There are no guidelines, just be kind and format your code!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
