package util

func Contains(items []Item, target string) bool {
	for _, item := range items {
		if str, ok := item.(string); ok && target == str {
			return true
		}
	}
	return false
}
