package shellwords

func Fuzz(data []byte) int {
	_, err := SplitPosix(string(data))
	if err != nil {
		return 0
	}

	_, err = SplitBatch(string(data))
	if err != nil {
		return 0
	}

	return 1
}
