# cryptipass

[![Go Report Card](https://goreportcard.com/badge/github.com/francescoalemanno/cryptipass)](https://goreportcard.com/report/github.com/francescoalemanno/cryptipass)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/francescoalemanno/cryptipass?status.svg)](https://pkg.go.dev/github.com/francescoalemanno/cryptipass)

**cryptipass** is a flexible, high-entropy passphrase generator that creates secure, pronounceable passwords using a probabilistic model. It's designed for security-conscious developers who need memorable yet strong passphrases.

---

## Features

- **Pronounceable Passwords**: Generates words based on real-world token patterns, making them easy to remember.
- **Highly Customizable**: Define your own word list or use pre-defined patterns like symbols, numbers, and mixed-case letters.
- **Secure Randomness**: Uses cryptographic-grade randomness (`crypto/rand`) for generating passphrases.
- **Entropy Analysis**: Built-in entropy calculations and certification to ensure high randomness and strength.
- **Pattern-Based Generation**: Control password structure using customizable patterns (e.g., words, digits, symbols).
  
---

## Installation

To install `cryptipass`, use `go get`:

```bash
go get github.com/francescoalemanno/cryptipass
```

Then, import it into your project:

```go
import "github.com/francescoalemanno/cryptipass"
```

NOTE: **We also have a [CLI](cmd/genpw) available for non-library uses.**

---

## Quick Start

Here's how to generate a passphrase using the default word style:

```go
package main

import (
	"fmt"
	"github.com/francescoalemanno/cryptipass"
)

func main() {
	// Create a new cryptipass generator
	gen := cryptipass.NewInstance()

	// Generate a 4-word passphrase
	passphrase, entropy := gen.GenPassphrase(4) 

	fmt.Println("Passphrase:", passphrase) //e.g. netica.peroundl.opantmene.symnals
	fmt.Println("Entropy:", entropy)
}
```

Want more control over the pattern? Use `GenFromPattern`:

```go
// Generate a password with pattern: Word-Number-Symbol
pass, entropy := gen.GenFromPattern("w-d-s") // eg. opantmene-4-%
fmt.Println("Generated Password:", pass)
```

Possible patterns are formed by combining:
- 'w' lowercase word, 'W' for uppercase word.
- 'c' a lowercase character, 'C' a uppercase character.
- 's' symbol, 'd' digit.
  
other symbols are interpolated in the final password and to interpolate one of the reserved symbols use escaping with "\".

---

## Custom Word Lists

You can customize the word style by creating a new instance from your own token set:

```go
myTokens := []string{"alpha", "bravo", "charlie", "delta"}
gen := cryptipass.NewCustomInstance(myTokens, 1) //instead of 1, try 2,3,4 to see the tradeoff between fidelity to the wordlist and entropy gain.

pass, entropy := gen.GenPassphrase(3)
fmt.Println("Custom Passphrase:", pass) //e.g. alphar.bravo.delta
fmt.Println("Entropy:", entropy)
```

---

## Documentation

Full API documentation is available at [GoDoc](https://pkg.go.dev/github.com/francescoalemanno/cryptipass).

---

## License

`cryptipass` is licensed under the MIT License. See [LICENSE](LICENSE) for details.

---

## Contributing

Contributions, issues, and feature requests are welcome! Feel free to check out [issues](https://github.com/francescoalemanno/cryptipass/issues) or open a pull request.

---

**cryptipass** – Secure, flexible, and pronounceable passphrases for your Go applications.
