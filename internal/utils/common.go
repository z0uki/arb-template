package utils

import "time"

func UTCToLocal(utc string) time.Time {
	//2022-12-01T16:25:22.000000+00:0
	ti, err := time.Parse(time.RFC3339, utc)
	if err != nil {
		return time.Now()
	}
	return ti.In(time.Local)
}
