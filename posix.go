package shellwords

// SplitPosix splits a command string into words like a posix shell would
func SplitPosix(line string) ([]string, error) {
	p := parser{
		Input:            line,
		QuoteChars:       []rune{'\'', '"'},
		EscapeChar:       '\\',
		QuoteEscapeChars: []rune{'\\'},
		FieldSeperators:  []rune{'\n', '\t', ' '},
	}
	return p.Parse()
}
