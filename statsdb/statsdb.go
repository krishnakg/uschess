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
	PreRating  int    `json:"preRating"`
	PostRating int    `json:"postRating"`
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

type SectionResult struct {
	SectionID  string  `json:"sectionId,omitempty"`
	PlayerID   string  `json:"playerId,omitempty"`
	PlayerName string  `json:"playerName,omitempty"`
	RatingType string  `json:"ratingType,omitempty"`
	PreRating  int     `json:"preRating"`
	PostRating string  `json:"postRating"`
	Score      float64 `json:"score"`
}

type Game struct {
	GameID       int    `json:"gameId,omitempty"`
	SectionID    string `json:"sectionid,omitempty"`
	Round        int    `json:"round,omitempty"`
	Result       string `json:"result,omitempty"`
	Player1ID    int    `json:"player1Id,omitempty"`
	Player1Name  string `json:"player1Name,omitempty"`
	Player1Color uint8  `json:"player1Color,omitempty"`
	Player2ID    int    `json:"player2Id,omitempty"`
	Player2Name  string `json:"player2Name,omitempty"`
	Player2Color uint8  `json:"player2Color,omitempty"`
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

func (stats *StatsDB) GetSectionResults(sectionId string) (results []SectionResult, err error) {
	rows, err := stats.db.Query("select th.section_id, th.uscf_id, p.name, th.rating_type, th.pre_rating, th.post_rating, th.score "+
		"from tournament_history th, player p where p.id=th.uscf_id and th.rating_type='R' and th.section_id=?", sectionId)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var result SectionResult
		if err := rows.Scan(&result.SectionID, &result.PlayerID, &result.PlayerName,
			&result.RatingType, &result.PreRating, &result.PostRating, &result.Score); err != nil {
			utils.CheckErr(err)
			return results, err
		}
		results = append(results, result)
	}
	return
}

func (stats *StatsDB) GetGames(sectionId string) (games []Game, err error) {
	rows, err := stats.db.Query("select g.id, g.section_id, g.round, g.result, "+
		"g.player1, p1.name, g.player1_color, g.player2, p2.name, g.player2_color "+
		"from player p1, game g, player p2 "+
		"where g.section_id=? and  p1.id=g.player1 and  p2.id=g.player2 and g.event_id in "+
		"(select event_id from tournament_history where rating_type='R');", sectionId)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.GameID, &game.SectionID, &game.Round, &game.Result,
			&game.Player1ID, &game.Player1Name, &game.Player1Color,
			&game.Player2ID, &game.Player2Name, &game.Player2Color); err != nil {
			utils.CheckErr(err)
			return games, err
		}
		games = append(games, game)
	}
	return
}
