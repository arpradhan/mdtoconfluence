// Package mdtoconfluence provides functions for converting confluence
// wiki markup to markdown
package mdtoconfluence

import (
	"regexp"
	"strconv"
	"strings"
)

// ReplaceHeading returns a heading formatted in confluence from markdown
func ReplaceHeading(str string) string {
	re := regexp.MustCompile(`^h[0-6]\. `)
	matched := re.FindString(str)
	if matched == "" {
		return str
	}
	indexes := re.FindStringIndex(str)
	re = regexp.MustCompile(`[0-6]`)
	matched = re.FindString(str)
	headingNumber, _ := strconv.Atoi(matched)
	headingSlice := make([]string, 0)
	for i := 0; i < headingNumber; i++ {
		headingSlice = append(headingSlice, "#")
	}

	headingSlice = append(headingSlice, str[indexes[1]-1:len(str)])
	parsed := strings.Join(headingSlice, "")
	return parsed
}
