package handler

import "time"

var layout = "2006-01-02"

func stringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}
