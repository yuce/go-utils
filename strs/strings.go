package strs

// FirstNonEmpty returns the first non-empty string from the given strings.
func FirstNonEmpty(items ...string) string {
	for _, item := range items {
		if item != "" {
			return item
		}
	}
	return ""
}
