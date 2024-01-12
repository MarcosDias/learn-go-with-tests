package iteration

import "strings"

const REPEAT_COUNT = 5

func Repeat(character string, count int) string {
	return strings.Repeat(character, count)
}
