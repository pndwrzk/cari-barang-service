package utils

import (
	"strings"
)

func GenerateSlug(input string) string {
	// Convert the string to lowercase
	slug := strings.ToLower(input)

	// Remove non-alphanumeric characters (except for spaces and hyphens)
	slug = RemoveSpecialChars(slug)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove multiple hyphens (if any)
	slug = strings.Join(strings.Fields(slug), "-")

	// Trim hyphens at the beginning and end
	slug = strings.Trim(slug, "-")

	return slug
}
