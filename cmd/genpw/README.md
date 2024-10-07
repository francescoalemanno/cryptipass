# genpw

**genpw** is a command-line tool that generates secure, pronounceable passwords and passphrases using customizable patterns. Built on top of the `cryptipass` library, **genpw** provides flexible options for creating memorable, yet strong passphrases with high entropy, making it ideal for secure applications, authentication, and more.

## Features

- Generate secure, pronounceable passwords.
- Customizable patterns for letters, numbers, symbols, and word formats.
- Visual strength meter for each generated passphrase.
- Cryptographic random number generator for secure randomness.
- High entropy and randomness assurance.

---

## Installation

To install **genpw**, use `go install`:

```bash
go install github.com/francescoalemanno/cryptipass/v2/cmd/genpw@latest
```

Ensure `$GOPATH/bin` is added to your `PATH` to run the command directly.

---

## Usage

You can generate a set of passwords with the desired pattern using the following command:

```bash
genpw -p [PATTERN] -n [NUMBER OF PASSWORDS (default 6)] -d [CHAIN DEPTH (default 2)]
```

### Pattern Options

- **w**: lowercase word.
- **W**: uppercase word.
- **c**: lowercase character.
- **C**: uppercase character.
- **d**: digit.
- **s**: symbol.

### Example

Generate 6 passphrases using a pattern of three words (with one capitalized), a two-digit number, and two lowercase characters:

```bash
genpw -p "W.w.w.ddcc" -n 6
```

Sample output:

```bash    
Passphrase                          Log10(Guesses)    Log2Entropy      Strength
                                                                      
Mortw.retainish.quater.66es            21.49             72.40       [===========.]
Defamito.repeac.stateryb.91he          23.75             79.90       [============]
Grazi.subsider.pravi.83de              20.14             67.92       [==========..]
Alishin.dumpedial.prayingin.45po       23.85             80.23       [============]
Atoryone.imputt.moodly.76op            19.66             66.32       [==========..]
Hurred.buffyeare.uphonetin.97co        23.22             78.12       [============]
```

### Parameters

- **`-p`**: Define the pattern for password generation (default: `W.w.w`).
- **`-n`**: Specify the number of passwords to generate (default: `6`).

---

## Password Strength

Each generated passphrase includes a strength meter, represented by a bar between `0` and `12` characters, where `=` indicates strength and `.` indicates relative weakness.

- **Log10(Guesses)**: The estimated number of guesses (in log10 scale) required to crack the password.
- **Log2Entropy**: The entropy (in bits) of the passphrase, reflecting its unpredictability.

---

## Customization

You can define your own patterns to create passphrases tailored to your security and usability needs. Patterns can be mixed and matched to create combinations of words, digits, symbols, and characters, offering high flexibility.

Examples:

- **Pattern**: `"wW.d.s"`
    - Output: lowercase word, uppercase word, digit, and symbol.

```bash
genpw -p "wW.d.s" -n 4
```

---

## License

**genpw** is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request for bug fixes, new features, or improvements.
