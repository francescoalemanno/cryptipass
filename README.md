# Cryptipass
NOTE: **We also have a [CLI](cmd/genpw) available for non-library uses.**
Cryptipass is a Go package designed to generate secure passphrases composed of human-readable words. The passphrases are generated with a focus on both security (through entropy) and usability by combining cryptographic randomness and customizable word generation strategies. 

## Features

- **Cryptographically secure randomization**: Uses `crypto/rand` for generating random data to seed the passphrase generation process, ensuring the highest level of security.
- **Customizable passphrase length**: The number of words in the passphrase can be controlled by the user.
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

The primary function of the `cryptipass` package is to generate secure passphrases. You can generate a new passphrase using the `NewPassphrase` function. You can specify how many words you want in the passphrase, and the function returns the passphrase and its total entropy.

Example:

```go
package main

import (
	"fmt"
	"cryptipass"
)

func main() {
	passphrase, entropy := cryptipass.NewPassphrase(5)
	fmt.Printf("Passphrase: %s\n", passphrase)
	fmt.Printf("Entropy: %.2f bits\n", entropy)
}
```

### Example Output:

```
Passphrase: jesside.flyperm.aunsis.dertsy
Entropy: 97.63 bits
```

### Word Generation

Internally, the package uses a series of functions to generate words of varying lengths. Each word contributes a certain amount of entropy, calculated during the generation process.

- `GenMixWord()`: Generates a random word of mixed length, returning both the word and its entropy.
- `GenWord(n int)`: Generates a word of exactly `n` characters.
- `PickLength()`: Picks a random length for a word.
- `PickNext()`: Generates the next part of a word based on the current string.

## Notes

- The package seeds the random number generator with `crypto/rand`, making it cryptographically secure. In scenarios where cryptographic security is not necessary and faster execution is preferred, the package also provides an alternative (commented out) `PCG` random number generator.
- The entropy provided in the output is a measure of how unpredictable the passphrase is. The higher the entropy, the more secure the passphrase is.

## Contributing

Contributions are welcome! If you encounter any issues or have feature suggestions, feel free to open an issue or a pull request on the GitHub repository.

## License

This project is licensed under the MIT License.

---

Happy coding and stay secure with Cryptipass!