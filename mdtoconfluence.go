// Package mdtoconfluence provides functions for converting confluence
// wiki markup to markdown
package mdtoconfluence

import (
	"fmt"
	"regexp"
	"strings"
)

// ReplaceStringHeading returns a heading formatted in confluence from markdown
func ReplaceStringHeading(str string) string {
	re := regexp.MustCompile(`(#{1,6}).+\n`)
	matches := re.FindAllStringSubmatch(str, -1)
	if matches == nil {
		return str
	}
	parsedSlice := make([]string, 0)
	for _, match := range matches {
		headingNumber := len(strings.Split(match[1], ""))
		heading := fmt.Sprintf("h%v.", headingNumber)
		re = regexp.MustCompile(`#{1,6}(.+\n)`)
		rest := re.FindStringSubmatch(match[0])[1]
		all := fmt.Sprintf(`%v%v`, heading, rest)
		parsedSlice = append(parsedSlice, all)

	}
	return strings.Join(parsedSlice, "")
}
