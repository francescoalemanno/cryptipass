# Cryptipass
NOTE: **We also have a [CLI](cmd/genpw) available for non-library uses.**
Cryptipass is a Go package designed to generate secure passphrases composed of human-readable words. The passphrases are generated with a focus on both security (through entropy) and usability by combining cryptographic randomness and customizable word generation strategies. 

## Features

- **Cryptographically secure randomization**: Uses `crypto/rand` for generating random data to seed the passphrase generation process, ensuring the highest level of security.
- **Customizable passphrase structure**: The number of words, symbols, digits in the passphrase can be controlled by the user.
- **Entropy calculation**: Provides an exact evaluation of the total entropy for the generated passphrase, helping users understand the strength of their passphrase.
- **Configurable word lengths**: Words within the passphrase can vary in length, ensuring better randomness and complexity.

## Installation

To use the `cryptipass` package in your project, you need to install it using Go's package management:

```bash
go get github.com/francescoalemanno/cryptipass
```

Then import it in your Go files:

```go
import "github.com/francescoalemanno/cryptipass"
```

## Usage

### Generate a Passphrase

The primary function of the `cryptipass` package is to generate secure passphrases. You can generate a new passphrase using the `GenPassphrase` function. You can specify how many words you want in the passphrase, and the function returns the passphrase and its total entropy.

Example:

```go
package main

import (
	"fmt"
	"github.com/francescoalemanno/cryptipass"
)

func main() {
	cp := cryptipass.NewInstance()
	passphrase, entropy := cp.GenPassphrase(5)
	fmt.Printf("Passphrase: %s\n", passphrase)
	fmt.Printf("Entropy: %.2f bits\n", entropy)
}
```

### Example Output:

```
Passphrase: ardram.iondbagro.anhambler.scheemous.chmedi
Entropy: 133.21 bits
```

### Generate a Password according to a pattern

You can generate a new password using the `GenFromPattern` function. You can specify how many words, symbols, digits you want in the passphrase by defining a pattern.

Possible patterns are formed by combining:
- 'w' lowercase word, 'W' for uppercase word.
- 'c' a lowercase character, 'C' a uppercase character.
- 's' symbol, 'd' digit.
- 
other symbols are interpolated in the final password and to interpolate one of the reserved symbols use escaping with "\".

The function returns the passphrase and its total entropy.

Example:

```go
package main

import (
	"fmt"
	"github.com/francescoalemanno/cryptipass"
)

func main() {
	cp := cryptipass.NewInstance()
	passphrase, entropy := cp.GenFromPattern("W-w.cccc.CCCC(ss)[20dd]")
	fmt.Printf("Passphrase: %s\n", passphrase)
	fmt.Printf("Entropy: %.2f bits\n", entropy)
}
```

### Example Output:

```
// pattern     W   -    w   .cccc.CCCC(ss)[20dd]
Passphrase: Storegu-dedudend.skin.EALR(=*)[2045]
Entropy: 96.41 bits
```

## Notes

- The package seeds the random number generator with `crypto/rand`, making it cryptographically secure.
- The entropy provided in the output is a measure of how unpredictable the passphrase is. The higher the entropy, the more secure the passphrase is.

## Contributing

Contributions are welcome! If you encounter any issues or have feature suggestions, feel free to open an issue or a pull request on the GitHub repository.

## License

This project is licensed under the MIT License.

---

Happy coding and stay secure with Cryptipass!