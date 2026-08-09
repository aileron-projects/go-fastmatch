package fastmatch_test

import (
	"path"
	"path/filepath"
	"testing"

	"github.com/aileron-projects/go-fastmatch"
	"github.com/aileron-projects/go-tester"
)

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

func TestBenches(t *testing.T) {
	matched, err := fastmatch.Match(matchPattern, matchStr)
	tester.AssertEqual(t, true, matched)
	tester.AssertEqual(t, nil, err)

	matched, err = path.Match(matchPattern, matchStr)
	tester.AssertEqual(t, true, matched)
	tester.AssertEqual(t, nil, err)

	matched, err = filepath.Match(matchPattern, matchStr)
	tester.AssertEqual(t, true, matched)
	tester.AssertEqual(t, nil, err)
}

func BenchmarkFastMatch(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		fastmatch.Match(matchPattern, matchStr)
	}
}

func BenchmarkPathMatch(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		path.Match(matchPattern, matchStr)
	}
}

func BenchmarkFilepathMatch(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		filepath.Match(matchPattern, matchStr)
	}
}
