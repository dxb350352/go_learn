package utils

import (
	"regexp"
	"strings"
)

var NumberReg = regexp.MustCompile(`\d+`)
func IsNumber(src string) bool {
	return NumberReg.MatchString(src)
}

func TrimLoopSuffixPrefix(s, fix string) string {
	if fix == "" {
		return s
	}
	for {
		if strings.HasSuffix(s, fix) {
			s = strings.TrimSuffix(s, fix)
		} else if strings.HasPrefix(s, fix) {
			s = strings.TrimPrefix(s, fix)
		} else {
			return s
		}
	}
	return ""
}