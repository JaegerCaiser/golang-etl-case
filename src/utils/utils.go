package utils

import (
	"regexp"
	"time"
)

var onlyNumbersRegex = regexp.MustCompile(`[^0-9]`)

//ConvertOnlyNumber - Remove não-digitos (números) da string
func ConvertOnlyNumber(s string) string {
	return onlyNumbersRegex.ReplaceAllString(s, "")
}

//ConvertStringToDate - Converte string para data
func ConvertStringToDate(s string) *time.Time {
	if s == "NULL" {
		return nil
	}

	format := "2006-01-02"

	s1, _ := time.Parse(format, s)

	return &s1
}
