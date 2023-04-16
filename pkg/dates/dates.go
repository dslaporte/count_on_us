package pkg_dates

import "time"

func StrToTime(str string) (time.Time, error) {
	return time.Parse(time.RFC3339, str)
}

func TimeToStr(date time.Time) string {
	return date.Format(time.RFC3339)
}
