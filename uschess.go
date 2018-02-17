package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	events := fetchEvents("2018-02-03")

	var stats StatsDB
	stats.open()
	defer stats.close()

	for _, event := range events {
		stats.saveEvent(event)
	}
}
