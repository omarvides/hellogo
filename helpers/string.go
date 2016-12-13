package helpers

// AddPrefix function adds the "prefix" word as prefix of the string received
func AddPrefix(prefix string, word string) string {
	return prefix + "-" + word
}

// AddSufix function adds a sufix to a word
func AddSufix(word string, sufix string) string {
	return word + "-" + sufix
}
