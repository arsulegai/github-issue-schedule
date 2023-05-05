package utils

import "time"

const layout = "2006-01-02"

func GetDate(date string) (time.Time, error) {
	// pattern match and get the date
	return time.Parse(layout, date)
}
