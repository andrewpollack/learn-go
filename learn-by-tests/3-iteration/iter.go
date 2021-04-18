package iter

func Repeat(character string, numRepeats int) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated += character
	}

	return repeated
}
