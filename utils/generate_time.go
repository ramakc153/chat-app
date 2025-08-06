package utils

import "time"

func Generate_time() time.Time {
	layout := "2025-04-17 19:02:24.807555+07"
	current_time := time.Now()
	formatted_time := current_time.Format(layout)
	parsed_layout, _ := time.Parse(layout, formatted_time)
	return parsed_layout
}
