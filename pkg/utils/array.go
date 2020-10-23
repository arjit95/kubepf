package utils

// IndexStr finds the index of a string in an array
func IndexStr(arr []string, str string) int {
	for i, v := range arr {
		if v == str {
			return i
		}
	}

	return -1
}

// FilterStr filters the array based on the callback function
func FilterStr(arr []string, filterFn func(str string) bool) []string {
	slice := make([]string, len(arr))

	for _, v := range arr {
		if filterFn(v) {
			slice = append(slice, v)
		}
	}

	return slice
}
