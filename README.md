# go-fastmatch

**Simple and fast text pattern matching library for Go.**

<div align="center">

[![GoDoc](https://godoc.org/github.com/aileron-projects/go-fastmatch?status.svg)](http://godoc.org/github.com/aileron-projects/go-fastmatch)
[![Test](https://github.com/aileron-projects/go-fastmatch/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/aileron-projects/go-fastmatch/actions/workflows/test.yaml?query=branch%3Amain)
[![License](https://img.shields.io/badge/License-Apache%202.0-yellow.svg)](./LICENSE)

[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/aileron-projects/go-fastmatch)
[![OpenSourceInsight](https://badgen.net/badge/open%2Fsource%2F/insight/cyan)](https://deps.dev/go/github.com%2Faileron-projects%2Fgo-fastmatch)
[![OSS Insight](https://badgen.net/badge/OSS/Insight/orange)](https://ossinsight.io/analyze/aileron-projects/go-fastmatch)

</div>

## Features

- Fast pattern matching
- Simple expression
- General pourpose

Why we need simple & fast pattern matcher?

- Standard [path.Match](https://pkg.go.dev/path#Match) is **specific** for path matching
- Standard [filepath.Match](https://pkg.go.dev/path/filepath#Match) is **specific** for file path matching
- Standard [regexp](https://pkg.go.dev/regexp) is for fully featured pattern matching but **slow**

## Tested Environments

Operating System:

- `Linux`: [ubuntu-latest](https://github.com/actions/runner-images)
- `Windows`: [windows-latest](https://github.com/actions/runner-images)
- `macOS`: [macos-latest](https://github.com/actions/runner-images)

Architecture (Using QEMU on linux):

- x86: `amd64`, `386`
- arm: `arm/v5`, `arm/v6`, `arm/v7`, `arm64`
- risc: `riscv64`, `loong64`
- ppc: `ppc64`, `ppc64le`
- mips: `mips`, `mips64`, `mips64le`, `mipsle`
- ibm: `s390x`

## Release Cycle

- Releases are made as needed.
- [Semantic Versioning](https://semver.org/) `vX.Y.Z` is used.

## License

[Apache-2.0](LICENSE)

## Usage

### Syntax

The syntax is very similar with standard package [path.Match](https://pkg.go.dev/path#Match) and [filepath.Match](https://pkg.go.dev/path/filepath#Match).

```txt
pattern:
    { term }
term:
    '*'         matches any sequence of characters
    '?'         matches any single character
    c           matches character c (c != '*', '?')
    '\\' c      matches character c
```

### Basic example

```go
pattern := `Hello *!`

fastmatch.Match(pattern, `Hello world !!`) // true
fastmatch.Match(pattern, `Hello Go !!`)    // true
fastmatch.Match(pattern, `Hello world ??`) // false
```

### Special character escape

Use `\` to escape special characters `*`, `?`.

```go
pattern := `Hello \*!`

fastmatch.Match(pattern, `Hello world !!`) // false
fastmatch.Match(pattern, `Hello *!`)       // true
```

## Benchmark

Used pattern and match string are:

```go
var matchStr = `
Go is
- An open-source programming language supported by Google.
- Easy to learn and great for teams.
- Built-in concurrency and a robust standard library.
- Large ecosystem of partners, communities, and tools.
`

var matchPattern = `
Go is
- An op*-so* p?o?r?m?i?g l*e su*orted by Google.
- Easy to le??? and great *** teams.
- Built-in c*o*n*c*u*r*r*e*n*c*y and a robust stand* lib*.
- Large ecosystem of partners, communities, and tools.*
`
```

Results (formatted):

```txt

                         iteration          speed     heap    allocation
                           -------  -------------   ------   -----------
BenchmarkFastMatch-8       1796689    658.6 ns/op   0 B/op   0 allocs/op ★
BenchmarkPathMatch-8        574101   2309.0 ns/op   0 B/op   0 allocs/op
BenchmarkFilepathMatch-8    722451   1913.0 ns/op   0 B/op   0 allocs/op
```
