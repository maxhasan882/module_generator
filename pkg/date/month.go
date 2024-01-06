package date_helper

import "time"

func MonthCount(st time.Time, en time.Time) int {
	cnt := 0
	for {
		cnt += 1
		st = st.AddDate(0, 1, 0)
		if st.Month() == en.Month() && st.Year() == en.Year() {
			cnt += 1
			break
		}
	}
	return cnt
}
