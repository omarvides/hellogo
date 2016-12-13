package helpers

// AddPrefix function adds the "prefix" word as prefix of the string received
func AddPrefix(prefix string, value string) string {
	return prefix + "-" + value
}
