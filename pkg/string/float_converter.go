package string_formatter

import (
	"strconv"
	"strings"
)

// FormatFloat this function first uses strconv.FormatFloat to convert the float value to a string with exactly 2 decimal places. Then, it uses strings.TrimSuffix to remove the trailing ".00" if it exists, making the formatted string cleaner
func FormatFloat(v float64) string {
	return strings.TrimSuffix(ConvertToBanglaNumeric(strconv.FormatFloat(v, 'f', 2, 64)), ".00")
}
