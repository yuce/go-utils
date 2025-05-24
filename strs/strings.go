package strs

import "strings"

// FirstNonEmpty returns the first non-empty string from the given strings.
func FirstNonEmpty(items ...string) string {
	for _, item := range items {
		if item != "" {
			return item
		}
	}
	return ""
}

// SplitByComma splits the text by comma.
// Each part is trimmed, empty parts are kept.
// If a string is all-whitespaces, then nil is returned.
func SplitByComma(text string) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}
	items := strings.Split(text, ",")
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}
	return items
}
