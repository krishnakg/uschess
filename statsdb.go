package main

import (
	"database/sql"
	"fmt"
)

type StatsDB struct {
	db *sql.DB
}

func (stats StatsDB) saveEvents(events []Event) {
	insertStatement, err := stats.db.Prepare("insert into event values (?, ?, ?, ?, ?, ?, ?)")
	checkErr(err)
	defer insertStatement.Close()

	for _, event := range events {
		_, err = insertStatement.Exec(event.id, event.sections, event.state, event.city, event.date, event.players, event.name)
		checkErr(err)
	}
}

func (stats *StatsDB) open() {
	var err error
	dbType, dbConnectionString := getDatabaseConnectionString()
	stats.db, err = sql.Open(dbType, dbConnectionString)
	checkErr(err)
}

func (stats *StatsDB) close() {
	stats.db.Close()
}

func (stats *StatsDB) insertEvent(event Event) {
	insertStatement, err := stats.db.Prepare("insert into event values (?, ?, ?, ?, ?, ?, ?)")
	checkErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(event.id, event.numSections, event.state, event.city, event.date, event.players, event.name)
	checkErr(err)
}

func (stats *StatsDB) insertSection(id string, name string, eventId string) {
	insertStatement, err := stats.db.Prepare("insert into section values (?, ?, ?)")
	checkErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(id, name, eventId)
	checkErr(err)
}

func (stats *StatsDB) insertGame(eventId string, sectionId string, round int, player1 int, player1Color Color, player2 int, player2Color Color, result string) {
	insertStatement, err := stats.db.Prepare("insert into game " +
		"(event_id, section_id, round, player1, player1_color, player2, player2_color, result) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?)")
	checkErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(eventId, sectionId, round, player1, player1Color, player2, player2Color, result)
	checkErr(err)
}

func (stats *StatsDB) insertTournamentHistory(uscfId int, eventId string, sectionId string, ratingType string, preRating int, postRating int, score float64) {
	insertStatement, err := stats.db.Prepare("insert into tournament_history " +
		"(uscf_id, event_id, section_id, rating_type, pre_rating, post_rating, score) " +
		"values (?, ?, ?, ?, ?, ?, ?)")
	checkErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(uscfId, eventId, sectionId, ratingType, preRating, postRating, score)
	checkErr(err)
}

func (stats *StatsDB) insertPlayer(uscfId int, name string, state string) {
	insertStatement, err := stats.db.Prepare("insert into player " +
		"(id, name, state) " +
		"values (?, ?, ?) " +
		"on duplicate key " +
		"update name=?, state=?")
	checkErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(uscfId, name, state, name, state)
	checkErr(err)
}

func (stats *StatsDB) saveEvent(event Event) {
	stats.insertEvent(event)
	for id, section := range event.sections {
		stats.insertSection(id, section.name, event.id)
		// Used to make a map of player position to USCF Id so as to use later.
		uscfID := make(map[int]int)
		for _, entry := range section.entries {
			// We should write the player table first as it does not have foriegn keys and also
			// because the player id will be used as the foriegn key in other tables.
			stats.insertPlayer(entry.id, entry.name, entry.state)
			for _, ratingChange := range entry.change {
				// Currently not storing the number of games for provisional players
				preRating, _ := parseRating(ratingChange.pre)
				postRating, _ := parseRating(ratingChange.post)
				fmt.Println(entry.id, event.id, section.id, ratingChange.ratingType, preRating, postRating, entry.score)
				stats.insertTournamentHistory(entry.id, event.id, section.id, ratingChange.ratingType, preRating, postRating, entry.score)
			}
			uscfID[entry.position] = entry.id
		}
		for _, entry := range section.entries {
			for i, game := range entry.games {
				if game.player2 > game.player1 {
					stats.insertGame(event.id, section.id, i+1, uscfID[game.player1], game.player1Color, uscfID[game.player2], game.player2Color, game.result)
				}
			}
		}
	}
}
