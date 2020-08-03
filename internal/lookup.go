package internal

import (
	"strings"
)

/*
Normalize strips special characters from strings and forces them to
lower-case to make the likelihood of a correct look-up greater.
*/
func Normalize(v string) string {
	sb := strings.Builder{}
	for _, c := range v {
		c = c | 0x20
		if c >= 'a' && c <= 'z' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func Unique(v []string) bool {
	set := map[string]bool{}
	for _, s := range v {
		if set[s] {
			return false
		}
		set[s] = true
	}

	return true
}
