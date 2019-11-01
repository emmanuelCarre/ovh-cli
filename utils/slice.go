package utils

// StringSliceContains return true if string is contained by a Slice
func StringSliceContains(sl []string, str string) bool {
	for _, item := range sl {
		if str == item {
			return true
		}
	}
	return false
}
