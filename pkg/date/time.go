package date_helper

import "time"

// TimeBeginningOfMonth return the begin of the month of t
func TimeBeginningOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 23, 59, 59, 59, t.Location())
}

// TimeEndOfMonth return the end of the month of t
func TimeEndOfMonth(t time.Time) time.Time {
	return TimeBeginningOfMonth(t).AddDate(0, 1, -1)
}

func TimeBeginningOfDay(t time.Time) time.Time {
	return t.Truncate(time.Hour * 24)
}

func TimeEndOfDay(t time.Time) time.Time {
	return TimeBeginningOfDay(t).AddDate(0, 0, 1)
}

func CurrentTimeBD() time.Time {
	now := time.Now().UTC().Add(time.Hour * 6)
	return now.Round(time.Second)
}
