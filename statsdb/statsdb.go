package statsdb

import (
	"database/sql"
	"fmt"
	"uschess/utils"
)

type StatsDB struct {
	db *sql.DB
}

type Player struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
}

type Section struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EventPerformance struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	SectionID  string `json:"sectionId,omitempty"`
	UscfID     int    `json:"uscfId,omitempty"`
	PreRating  int    `json:"preRating,omitempty"`
	PostRating int    `json:"postRating,omitempty"`
	RatingType string `json:"ratingType,omitempty"`
}

type Tournament struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	State    string `json:"state,omitempty"`
	City     string `json:"city,omitempty"`
	Players  int    `json:"players,omitempty"`
	Sections int    `json:"sections,omitempty"`
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

func (stats *StatsDB) GetPlayer(uscfId int) (player Player, err error) {
	err = stats.db.QueryRow("select name, state from player where id=?", uscfId).Scan(&player.Name, &player.State)
	player.ID = uscfId
	return
}

func (stats *StatsDB) GetTournament(tournamentId string) (tournament Tournament, err error) {
	err = stats.db.QueryRow("select name, state, city, players, sections from event where id=?", tournamentId).Scan(
		&tournament.Name, &tournament.State, &tournament.City, &tournament.Players, &tournament.Sections)
	tournament.ID = tournamentId
	return
}

func (stats *StatsDB) GetSectionInfo(tournamentId string) (sections []Section, err error) {
	rows, err := stats.db.Query("select id, name from section where event_id=?", tournamentId)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var section Section
		if err := rows.Scan(&section.ID, &section.Name); err != nil {
			utils.CheckErr(err)
			return sections, err
		}
		sections = append(sections, section)
	}
	return
}

func (stats *StatsDB) GetPlayerSearchResult(query string) (players []Player, err error) {
	query = fmt.Sprintf("%s%%", query)
	rows, err := stats.db.Query("select id, name, state from player where name like ? limit 10", query)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var player Player
		if err := rows.Scan(&player.ID, &player.Name, &player.State); err != nil {
			utils.CheckErr(err)
			return players, err
		}
		players = append(players, player)
	}
	return
}

func (stats *StatsDB) GetEvents(uscfId int) (performances []EventPerformance, err error) {
	rows, err := stats.db.Query("select e.id, e.name, th.section_id, th.uscf_id, th.pre_rating, th.post_rating, th.rating_type "+
		"from event e, tournament_history th where th.event_id=e.id and th.rating_type='R' and th.uscf_id=? order by e.id desc", uscfId)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var performance EventPerformance
		if err := rows.Scan(&performance.ID, &performance.Name, &performance.SectionID, &performance.UscfID,
			&performance.PreRating, &performance.PostRating, &performance.RatingType); err != nil {
			utils.CheckErr(err)
			return performances, err
		}
		performances = append(performances, performance)
	}
	return
}
