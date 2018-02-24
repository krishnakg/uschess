package main

import (
	"fmt"
	"uschess/statsdb"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	events := fetchEvents("2018-02-04")

	var stats statsdb.StatsDB
	stats.Open()
	defer stats.Close()

	for _, event := range events {
		saveEvent(stats, event)
	}
}

func saveEvent(stats statsdb.StatsDB, event Event) {
	stats.InsertEvent(event.id, event.numSections, event.state, event.city, event.date, event.players, event.name)
	for id, section := range event.sections {
		stats.InsertSection(id, section.name, event.id)
		// Used to make a map of player position to USCF Id so as to use later.
		uscfID := make(map[int]int)
		for _, entry := range section.entries {
			// We should write the player table first as it does not have foriegn keys and also
			// because the player id will be used as the foriegn key in other tables.
			stats.InsertPlayer(entry.id, entry.name, entry.state)
			for _, ratingChange := range entry.change {
				// Currently not storing the number of games for provisional players
				preRating, _ := parseRating(ratingChange.pre)
				postRating, _ := parseRating(ratingChange.post)
				fmt.Println(entry.id, event.id, section.id, ratingChange.ratingType, preRating, postRating, entry.score)
				stats.InsertTournamentHistory(entry.id, event.id, section.id, ratingChange.ratingType, preRating, postRating, entry.score)
			}
			uscfID[entry.position] = entry.id
		}
		for _, entry := range section.entries {
			for i, game := range entry.games {
				if game.player2 > game.player1 {
					stats.InsertGame(event.id, section.id, i+1, uscfID[game.player1], int(game.player1Color), uscfID[game.player2], int(game.player2Color), game.result)
				}
			}
		}
	}
}
