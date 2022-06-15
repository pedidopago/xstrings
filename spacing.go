package xstrings

import "regexp"

var (
	regexLeadingSpaces  = regexp.MustCompile(`^\s+`)
	regexTrailingSpaces = regexp.MustCompile(`\s+$`)
)

func TrimLeadingSpaces(v string) string {
	return regexLeadingSpaces.ReplaceAllString(v, "")
}

func TrimTrailingSpaces(v string) string {
	return regexTrailingSpaces.ReplaceAllString(v, "")
}

func TrimLeadingAndTrailingSpaces(v string) string {
	return TrimTrailingSpaces(TrimLeadingSpaces(v))
}
