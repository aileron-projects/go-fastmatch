package fastmatch

import (
	"errors"
)

// ErrBadPattern tells that the pattern is malformed.
var ErrBadPattern = errors.New("go-fastmatch/fastmatch: bad pattern")

// Match reports whether the str matches the pattern.
// Match is simpler and faster match function
// similar to [path.Match] or [path/filepath.Match].
// The only error returned is [ErrBadPattern] when the given pattern
// has single "\\" at the end of its pattern.
//
// The pattern syntax is:
//
//	pattern:
//		{ term }
//	term:
//		'*'         matches any sequence of characters
//		'?'         matches any single character
//		c           matches character c (c != '*', '?')
//		'\\' c      matches character c
func Match(pattern, str string) (bool, error) {
	if countByteRight(pattern, '\\')%2 == 1 {
		return false, ErrBadPattern // The only error pattern.
	}

	var star bool
	var chunk string
Loop:
	for len(pattern) > 0 {
		star, chunk, pattern = scanChunk(pattern)
		if star {
			if chunk == "" {
				return true, nil
			}
			for i := 0; i <= len(str); i++ {
				if ok, rest := matchChunk(chunk, str[i:]); ok {
					if len(pattern) == 0 && len(rest) > 0 {
						continue
					}
					str = rest
					continue Loop
				}
			}
			return false, nil
		} else {
			ok, rest := matchChunk(chunk, str)
			if !ok {
				return false, nil
			}
			str = rest
		}
	}
	return str == "", nil
}

// scanChunk scans the given pattern.
// scanChunk returns true if the given pattern has
// one or more "*" as its prefix.
// The chunk syntax is:
//
//	chunk:
//		{ term }
//	term:
//		'?'         matches any single character
//		c           matches character c (c != '*', '?')
//		'\\' c      matches character c
//
// Some examples of input/output are:
//
//	?        : false, "?", ""
//	?foo     : false, "?", "foo"
//	?foo*bar : false, "?", "foo*bar"
//	?foo?bar : false, "?", "foo?bar"
//	*        : true, "", ""
//	*foo     : true, "foo", ""
//	*foo*bar : true, "foo", "*bar"
//	*foo?bar : true, "foo?bar", ""
//	foo      : false, "foo", ""
//	foo*     : false, "foo", "*"
//	foo?     : false, "foo?", ""
//	foo*bar  : false, "foo", "*bar"
//	foo?bar  : false, "foo?bar", ""
func scanChunk(pattern string) (star bool, chunk, rest string) {
	escape := false
	for len(pattern) > 0 {
		if pattern[0] == '*' {
			star = true
			pattern = pattern[1:]
			continue
		}
		break
	}
	for i := 0; i < len(pattern); i++ {
		if escape {
			escape = false
			continue
		}
		switch pattern[i] {
		case '*':
			return star, pattern[:i], pattern[i:]
		case '\\':
			escape = true
		}
	}
	return star, pattern, ""
}

// matchChunk reports whether the prefix of the target string
// matches the pattern.
// If the given pattern has single '\\' at the end of the pattern,
// matchChunk returns false (pattern "str\\" does not match "str\\").
// The chunk syntax is:
//
//	chunk:
//		{ term }
//	term:
//		'?'         matches any single character
//		c           matches character c (c != '*', '?')
//		'\\' c      matches character c
func matchChunk(chunk, target string) (ok bool, rest string) {
	escape := false
	for len(chunk) > 0 {
		if len(target) == 0 {
			return false, ""
		}
		switch chunk[0] {
		case '?':
			if escape {
				escape = false
				break
			}
			chunk = chunk[1:]
			target = target[1:]
			continue
		case '\\':
			if escape {
				escape = false
				break
			}
			escape = true
			chunk = chunk[1:]
			continue
		}
		if chunk[0] == target[0] {
			chunk = chunk[1:]
			target = target[1:]
			continue
		}
		return false, target
	}
	return true, target
}

// CountBytecountByteRightRight counts the number of b from the
// right side of s until non b found.
// For example, countByteRight("abcc",'c') returns 2.
func countByteRight(s string, b byte) int {
	n := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == b {
			n += 1
		} else {
			break
		}
	}
	return n
}
