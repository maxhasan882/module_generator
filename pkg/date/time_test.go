package date_helper

import (
	"fmt"
	"testing"
)

func TestTimeBeginningOfDay(t *testing.T) {
	now := CurrentTimeBD()

	beg := TimeBeginningOfDay(now)
	fmt.Println(beg)
}
