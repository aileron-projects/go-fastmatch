<!-- markdownlint-disable MD033 MD041 -->

<div align="center">

[![Release](https://img.shields.io/github/v/release/aileron-projects/go-fastmatch?sort=semver)](https://github.com/aileron-projects/go-fastmatch/releases)
[![Reference](https://pkg.go.dev/badge/github.com/aileron-projects/go-fastmatch.svg)](https://pkg.go.dev/github.com/aileron-projects/go-fastmatch)
[![DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/aileron-projects/go-fastmatch)
[![Test](https://github.com/aileron-projects/go-fastmatch/actions/workflows/test.yaml/badge.svg)](https://github.com/aileron-projects/go/actions/workflows/test.yaml)

[![Insights](https://badgen.net/badge/Insights/open%2Fsource%2Finsights/cyan)](https://deps.dev/go/github.com%2Faileron-projects%2Fgo-fastmatch)
[![Insights](https://badgen.net/badge/Insights/OSS%2FInsight/orange)](https://ossinsight.io/analyze/aileron-projects/go-fastmatch)

</div>

# go-fastmatch

**Simple and fast text pattern matching library for Go.**

## Features

- Fast pattern matching
- Simple expression
- General pourpose

Why we need simple & fast pattern matcher?

- Standard [path.Match](https://pkg.go.dev/path#Match) is **specific** for path matching
- Standard [filepath.Match](https://pkg.go.dev/path/filepath#Match) is **specific** for file path matching
- Standard [regexp](https://pkg.go.dev/regexp) is for fully featured pattern matching but **slow**

## Usages

### Syntax

The syntax is very similar with standard [path.Match](https://pkg.go.dev/path#Match) and [filepath.Match](https://pkg.go.dev/path/filepath#Match).

```txt
pattern:
    { term }
term:
    '*'         matches any sequence of characters
    '?'         matches any single character
    c           matches character c (c != '*', '?', '\\')
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

Use `\` to escape special characters `*`, `?` and `\`.

```go
pattern := `Hello \*!`

fastmatch.Match(pattern, `Hello world !!`) // false
fastmatch.Match(pattern, `Hello *!`)       // true
```

## Docs & Examples

- GoDoc: <https://pkg.go.dev/github.com/aileron-projects/go-fastmatch>
- Examples: [example_test.go](./example_test.go)

## Benchmarks

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

## References

- [path#Match](https://pkg.go.dev/path#Match)
- [path/filepath#Match](https://pkg.go.dev/path/filepath#Match)
