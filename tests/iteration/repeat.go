package iteration

func Repeat(character string, repCount int) string {
	var repeated string
	repeatCount := repCount
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
