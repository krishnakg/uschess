package statsdb

import (
	"database/sql"
	"uschess/utils"
)

type StatsDB struct {
	db *sql.DB
}

func (stats *StatsDB) Open() {
	var err error
	dbType, dbConnectionString := utils.GetDatabaseConnectionString()
	stats.db, err = sql.Open(dbType, dbConnectionString)
	utils.CheckErr(err)
}

func (stats *StatsDB) Close() {
	stats.db.Close()
}

func (stats *StatsDB) InsertEvent(id string, numSections int, state string, city string, date string, players int, name string) {
	insertStatement, err := stats.db.Prepare("insert into event values (?, ?, ?, ?, ?, ?, ?)")
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(id, numSections, state, city, date, players, name)
	utils.CheckErr(err)
}

func (stats *StatsDB) InsertSection(id string, name string, eventId string) {
	insertStatement, err := stats.db.Prepare("insert into section values (?, ?, ?)")
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(id, name, eventId)
	utils.CheckErr(err)
}

func (stats *StatsDB) InsertGame(eventId string, sectionId string, round int, player1 int, player1Color int, player2 int, player2Color int, result string) {
	insertStatement, err := stats.db.Prepare("insert into game " +
		"(event_id, section_id, round, player1, player1_color, player2, player2_color, result) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?)")
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(eventId, sectionId, round, player1, player1Color, player2, player2Color, result)
	utils.CheckErr(err)
}

func (stats *StatsDB) InsertTournamentHistory(uscfId int, eventId string, sectionId string, ratingType string, preRating int, postRating int, score float64) {
	insertStatement, err := stats.db.Prepare("insert into tournament_history " +
		"(uscf_id, event_id, section_id, rating_type, pre_rating, post_rating, score) " +
		"values (?, ?, ?, ?, ?, ?, ?)")
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(uscfId, eventId, sectionId, ratingType, preRating, postRating, score)
	utils.CheckErr(err)
}

func (stats *StatsDB) InsertPlayer(uscfId int, name string, state string) {
	insertStatement, err := stats.db.Prepare("insert into player " +
		"(id, name, state) " +
		"values (?, ?, ?) " +
		"on duplicate key " +
		"update name=?, state=?")
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(uscfId, name, state, name, state)
	utils.CheckErr(err)
}
