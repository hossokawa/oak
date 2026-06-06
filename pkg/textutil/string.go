package textutil

import (
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	if s == "" {
		return ""
	}

	ss := strings.Split(s, " ")

	var capitalized []string
	if len(ss) >= 2 {
		for _, str := range ss {
			rs := []rune(str)
			rs[0] = unicode.ToUpper(rs[0])
			capitalized = append(capitalized, string(rs))
		}
		return strings.Join(capitalized, " ")
	}

	rs := []rune(ss[0])
	rs[0] = unicode.ToUpper(rs[0])

	return string(rs)
}

func Normalize(s string) string {
	if s == "" {
		return ""
	}

	s = strings.ReplaceAll(s, "-", " ")

	return s
}
