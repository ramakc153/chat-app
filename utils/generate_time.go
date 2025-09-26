package utils

import "time"

func Generate_time() time.Time {
	layout := "2006-01-02 15:04:05.000000-07"
	current_time := time.Now()
	formatted_time := current_time.Format(layout)
	parsed_layout, _ := time.Parse(layout, formatted_time)
	return parsed_layout
}
