package utils

import (
	"regexp"
	"strconv"
)

func RemoveSpecialChars(str string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9\\s-]")
	return re.ReplaceAllString(str, "")
}

func StrToUint(str string) (uint, error) {
	uintID, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err // Return 0 instead of nil
	}
	return uint(uintID), nil
}
