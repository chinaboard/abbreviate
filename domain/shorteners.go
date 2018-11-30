package domain

import (
	"strings"
	"unicode"
)

func findParts(original string) []string {
	parts := []string{}

	set := ""
	for _, ch := range []rune(original) {
		if !unicode.IsLetter(ch) {
			if set != "" {
				parts = append(parts, set)
			}
			parts = append(parts, string(ch))
			set = ""
		} else if unicode.IsUpper(ch) {
			if set != "" {
				parts = append(parts, set)
			}
			set = string(ch)
		} else if !unicode.IsLetter(ch) {
			if set != "" {
				parts = append(parts, set)
			}

			parts = append(parts, string(ch))
			set = ""
		} else {
			set += string(ch)
		}
	}
	if set != "" {
		parts = append(parts, set)
	}

	return parts
}

func sumLen(s []string) int {
	l := 0

	for _, p := range s {
		l = len(p)
	}
	return l
}

// Shortener represents an algorithm that can be used to shorten a string
// by substituting words for abbreviations
type Shortener func(matcher Matcher, original string, max int) string

// ShortenFromBack discovers words using camel case and non letter characters,
// starting from the back until the string has less than 'max' characters
// or it can't shorten any more.
func ShortenFromBack(matcher *Matcher, original string, max int) string {
	if len(original) < max {
		return original
	}

	parts := findParts(original)
	for pos := len(parts) - 1; pos >= 0 && sumLen(parts) > max; pos-- {
		str := parts[pos]
		abbr := matcher.Match(strings.ToLower(str))
		if isTitleCase(str) {
			abbr = makeTitle(abbr)
		}
		parts[pos] = abbr
	}

	return strings.Join(parts, "")
}

func isTitleCase(str string) bool {
	ch := first(str)
	return unicode.IsUpper(ch)
}

func makeTitle(str string) string {
	if str == "" {
		return ""
	}

	ch := first(str)
	ch = unicode.ToUpper(ch)
	result := string(ch)
	if len(str) > 1 {
		result += str[1:]
	}

	return result
}

func first(str string) rune {
	if str == "" {
		return rune(0)
	}

	return []rune(str)[0]
}

func lastChar(str string) (string, rune) {
	l := len(str)

	switch l {
	case 0:
		return "", rune(0)

	case 1:
		return "", []rune(str)[0]
	}

	return str[0 : l-1], []rune(str)[l-1:][0]
}
