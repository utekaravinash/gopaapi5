package api

// existsInStrings checks if a string exists in a slice of strings
func existsInStrings(value string, values []string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}
