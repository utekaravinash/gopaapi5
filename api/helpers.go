package api

func existsInStrings(value string, values []string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}
