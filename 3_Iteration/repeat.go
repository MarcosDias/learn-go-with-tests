package iteration

const REPEAT_COUNT = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < REPEAT_COUNT; i++ {
		repeated += character
	}
	return repeated
}
