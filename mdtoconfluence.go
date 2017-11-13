// Package mdtoconfluence provides functions for converting confluence
// wiki markup to markdown
package mdtoconfluence

import (
	"fmt"
	"regexp"
	"strings"
)

// ReplaceStringHeading returns a confluence formatted heading from markdown
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

// ReplaceStringBulletList returns a confluence formatted list from markdown
func ReplaceStringBulletList(str string) string {
	lines := strings.Split(str, "\n")
	newLines := make([]string, 0)
	for _, line := range lines {
		re := regexp.MustCompile(`^(\*|-)(.+)`)
		match := re.FindStringSubmatch(line)
		var newLine string
		if match == nil {
			newLine = line
		} else {
			rest := match[2]
			newLine = fmt.Sprintf("%s%s", "*", rest)
		}
		newLines = append(newLines, newLine)
	}
	return strings.Join(newLines, "\n")
}
