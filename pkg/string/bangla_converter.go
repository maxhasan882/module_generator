package string_formatter

import (
	"fmt"
	"strings"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func ConvertToBanglaNumeric[T Number](number T) string {
	regularDigits := "0123456789"
	banglaDigits := "০১২৩৪৫৬৭৮৯"

	regularToBangla := make(map[rune]rune)
	for i, r := range []rune(regularDigits) {
		regularToBangla[r] = []rune(banglaDigits)[i]
	}

	regularNumberStr := fmt.Sprintf("%v", number)
	var banglaNumberBuilder strings.Builder

	for _, char := range regularNumberStr {
		if banglaChar, found := regularToBangla[char]; found {
			banglaNumberBuilder.WriteRune(banglaChar)
		} else {
			banglaNumberBuilder.WriteRune(char)
		}
	}
	return banglaNumberBuilder.String()
}
