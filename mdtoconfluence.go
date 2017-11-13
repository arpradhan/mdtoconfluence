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

// ReplaceStringNestedBulletList returns a confluence formatted nested list from markdown
func ReplaceStringNestedBulletList(str string) string {
	lines := strings.Split(str, "\n")
	newLines := make([]string, 0)
	for i, line := range lines {
		re := regexp.MustCompile(` +(\*|-)`)
		match := re.FindString(line)
		var newLine string
		if match == "" {
			newLine = line
		} else {
			if i > 0 {
				lastLine := lines[i-1]
				re = regexp.MustCompile(`(\*|-) (.+)`)
				firstLetterIndex := re.FindStringSubmatchIndex(lastLine)[4]
				if len(match)-1 == firstLetterIndex {
					re = regexp.MustCompile(`(\*|-)`)
					nestedBulletIndex := re.FindStringIndex(line)[0]
					if nestedBulletIndex == firstLetterIndex {
						re = regexp.MustCompile(`( +).+`)
						bulletNumber := re.FindStringSubmatchIndex(line)[3]/2 + 1
						bulletSlice := make([]string, bulletNumber)
						for j := 0; j < bulletNumber; j++ {
							bulletSlice[j] = "*"
						}
						re = regexp.MustCompile(`(\*|-)(.+)`)
						rest := re.FindStringSubmatch(line)[2]
						newLine = fmt.Sprintf("%v%v", strings.Join(bulletSlice, ""), rest)
					}
				}

			}
		}
		newLines = append(newLines, newLine)
	}
	return strings.Join(newLines, "\n")
}
