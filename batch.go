package shellwords

// SplitBatch splits a command string into words like Windows CMD.EXE would
// See https://ss64.com/nt/syntax-esc.html
func SplitBatch(line string) ([]string, error) {
	p := parser{
		Input:            line,
		QuoteChars:       []rune{'\'', '"'},
		EscapeChar:       '^',
		QuoteEscapeChars: []rune{'^', '"'},
		FieldSeperators:  []rune{'\n', '\t', ' '},
	}
	return p.Parse()
}
