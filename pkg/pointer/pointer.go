package pointer

import "time"

func TimeP(t time.Time) *time.Time {
	return &t
}

func BoolP(b bool) *bool {
	return &b
}

func StringP(s string) *string {
	return &s
}
