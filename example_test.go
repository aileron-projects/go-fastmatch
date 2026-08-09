package fastmatch_test

import (
	"fmt"

	"github.com/aileron-projects/go-fastmatch"
)

func ExampleMatch() {
	pattern := `Hello *!`
	fmt.Println(fastmatch.Match(pattern, `Hello world !!`)) // true
	fmt.Println(fastmatch.Match(pattern, `Hello Go !!`))    // true
	fmt.Println(fastmatch.Match(pattern, `Hello world ??`)) // false
	// Output:
	// true <nil>
	// true <nil>
	// false <nil>
}

func ExampleMatch_escape() {
	pattern := `Hello \*!`
	fmt.Println(fastmatch.Match(pattern, `Hello world !!`)) // false
	fmt.Println(fastmatch.Match(pattern, `Hello *!`))       // true
	// Output:
	// false <nil>
	// true <nil>
}

func ExampleMatch_escapeAll() {
	fmt.Println(fastmatch.Match(`\*`, `*`))     // true
	fmt.Println(fastmatch.Match(`\*`, `hello`)) // false
	fmt.Println(fastmatch.Match(`\?`, `?`))     // true
	fmt.Println(fastmatch.Match(`\?`, `x`))     // false
	fmt.Println(fastmatch.Match(`\\`, `\`))     // true
	fmt.Println(fastmatch.Match(`\\`, `\\`))    // false
	// Output:
	// true <nil>
	// false <nil>
	// true <nil>
	// false <nil>
	// true <nil>
	// false <nil>
}
