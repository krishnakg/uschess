package statsdb

import (
	"database/sql"
	"fmt"
	"strings"
	"uschess/utils"
)

// StatsDB is a wrapper around a database connection.
type StatsDB struct {
	db *sql.DB
}

// Player describes all the information about a player.
type Player struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	State  string `json:"state,omitempty"`
	FideID int    `json:"fideId,omitempty"`
}

// Section describes all the information about a section.
type Section struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// EventPerformance describes the performance of a single player in a section.
// TODO: This looks very similar to the SectionResult datastructure. Look into
// merging these.
type EventPerformance struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	SectionID  string `json:"sectionId,omitempty"`
	UscfID     int    `json:"uscfId,omitempty"`
	PreRating  int    `json:"preRating"`
	PostRating int    `json:"postRating"`
	RatingType string `json:"ratingType,omitempty"`
}

// Tournament describes information about a tournament.
type Tournament struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	State    string `json:"state,omitempty"`
	City     string `json:"city,omitempty"`
	Players  int    `json:"players,omitempty"`
	Sections int    `json:"sections,omitempty"`
}

// SectionResult describes the performance of a single player in a section.
type SectionResult struct {
	SectionID  string  `json:"sectionId,omitempty"`
	PlayerID   string  `json:"playerId,omitempty"`
	PlayerName string  `json:"playerName,omitempty"`
	RatingType string  `json:"ratingType,omitempty"`
	PreRating  int     `json:"preRating"`
	PostRating string  `json:"postRating"`
	Score      float64 `json:"score"`
}

// Game describes all the information about a specific game.
type Game struct {
	GameID       int    `json:"gameId,omitempty"`
	EventName    string `json:"eventName,omitempty"`
	SectionID    string `json:"sectionId,omitempty"`
	Round        int    `json:"round,omitempty"`
	Result       string `json:"result,omitempty"`
	Player1ID    int    `json:"player1Id,omitempty"`
	Player1Name  string `json:"player1Name,omitempty"`
	Player1Color uint8  `json:"player1Color,omitempty"`
	Player2ID    int    `json:"player2Id,omitempty"`
	Player2Name  string `json:"player2Name,omitempty"`
	Player2Color uint8  `json:"player2Color,omitempty"`
}

// Open opens a database connection.
func (stats *StatsDB) Open() {
	var err error
	dbType, dbConnectionString := utils.GetDatabaseConnectionString()
	stats.db, err = sql.Open(dbType, dbConnectionString)
	utils.CheckErr(err)
}

// Close closes the database connection.
func (stats *StatsDB) Close() {
	stats.db.Close()
}

// InsertEvent inserts a row into the event table.
func (stats *StatsDB) InsertEvent(id string, numSections int, state string, city string, date string, players int, name string) {
	insertStatement, err := stats.db.Prepare(queryInsertEvent)
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(id, numSections, state, city, date, players, name)
	utils.CheckErr(err)
}

// InsertSection inserts a row into the section table.
func (stats *StatsDB) InsertSection(id string, name string, eventID string) {
	insertStatement, err := stats.db.Prepare(queryInsertSection)
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(id, name, eventID)
	utils.CheckErr(err)
}

// InsertGame inserts a row into the game table.
func (stats *StatsDB) InsertGame(eventID string, sectionID string, round int, player1 int, player1Color int, player2 int, player2Color int, result string) {
	insertStatement, err := stats.db.Prepare(queryInsertGame)
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(eventID, sectionID, round, player1, player1Color, player2, player2Color, result)
	utils.CheckErr(err)
}

// InsertTournamentHistory inserts a row into the tournament_history table.
func (stats *StatsDB) InsertTournamentHistory(uscfID int, eventID string, sectionID string, ratingType string, preRating int, postRating int, score float64) {
	insertStatement, err := stats.db.Prepare(queryInsertTournamentHistory)
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(uscfID, eventID, sectionID, ratingType, preRating, postRating, score)
	utils.CheckErr(err)
}

// InsertPlayer inserts a row into the player table. Different from other APIs here, this will overwrite name and state of the player
// if the player is already present in the table.
func (stats *StatsDB) InsertPlayer(uscfID int, name string, state string) {
	insertStatement, err := stats.db.Prepare(queryInsertPlayer)
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(uscfID, name, state, name, state)
	utils.CheckErr(err)
}

// InsertFideID inserts the fide_id value for the specified uscf_id.
func (stats *StatsDB) InsertFideID(uscfID int, fideID int) {
	insertStatement, err := stats.db.Prepare(queryInsertFideID)
	utils.CheckErr(err)
	defer insertStatement.Close()

	_, err = insertStatement.Exec(uscfID, fideID, fideID)
	utils.CheckErr(err)
}

// GetPlayer fetches the specified player information from the player table.
func (stats *StatsDB) GetPlayer(uscfID int) (player Player, err error) {
	err = stats.db.QueryRow(queryGetPlayer, uscfID).Scan(&player.Name, &player.State, &player.FideID)
	player.ID = uscfID
	return
}

// GetTournament fetches the tournament information for the specified tournament Id.
func (stats *StatsDB) GetTournament(tournamentID string) (tournament Tournament, err error) {
	err = stats.db.QueryRow(queryGetEvent, tournamentID).Scan(
		&tournament.Name, &tournament.State, &tournament.City, &tournament.Players, &tournament.Sections)
	tournament.ID = tournamentID
	return
}

// GetRecentTournaments fetches the specified number of recent tournaments.
func (stats *StatsDB) GetRecentTournaments(numTournaments int) (tournaments []Tournament, err error) {
	rows, err := stats.db.Query(queryGetRecentTournaments, numTournaments)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var tournament Tournament
		if err := rows.Scan(&tournament.ID, &tournament.Name, &tournament.City, &tournament.State,
			&tournament.Players); err != nil {
			utils.CheckErr(err)
			return tournaments, err
		}
		tournaments = append(tournaments, tournament)
	}
	return
}

// GetSectionInfo fetches the list of sections in a tournament and their associated info from the section table.
func (stats *StatsDB) GetSectionInfo(tournamentID string) (sections []Section, err error) {
	rows, err := stats.db.Query(queryGetSectionInfo, tournamentID)
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

// GetPlayerSearchResult fetches the top 10 search results on players for the specified query.
func (stats *StatsDB) GetPlayerSearchResult(query string) (players []Player, err error) {
	query = fmt.Sprintf("%s%%", query)
	rows, err := stats.db.Query(queryGetPlayerSearchResult, query)
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

// GetEvents fetches all tournaments played by the specified uscfId.
func (stats *StatsDB) GetEvents(uscfID int) (performances []EventPerformance, err error) {
	rows, err := stats.db.Query(queryGetEventsForPlayer, uscfID)
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

// GetSectionResults fetches list of all players who played a section and their results.
func (stats *StatsDB) GetSectionResults(sectionID string) (results []SectionResult, err error) {
	rows, err := stats.db.Query(queryGetSectionResults, sectionID)
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

// GetGames fetches information on all games played in a given section.
func (stats *StatsDB) GetGames(sectionID string) (games []Game, err error) {
	rows, err := stats.db.Query(queryGetAllGamesInSection, sectionID)
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

// DeleteEvents deletes all events in the database for the particular date. The date format should
// be YYYY-MM-DD.
func (stats *StatsDB) DeleteEvents(date string) {
	// Event Id is of the form YYYYMMDD... So we need to remove the "-" seperator in the input date.
	query := fmt.Sprintf("%s%%", strings.Replace(date, "-", "", -1))
	deleteStatement, err := stats.db.Prepare(queryDeleteEvent)
	utils.CheckErr(err)
	defer deleteStatement.Close()

	_, err = deleteStatement.Exec(query)
	utils.CheckErr(err)
}

// GetMutualGames fetches all games played between player1 and player2.
func (stats *StatsDB) GetMutualGames(player1ID int, player2ID int) (games []Game, err error) {
	rows, err := stats.db.Query(queryGetMutualGames, player1ID, player2ID, player2ID, player1ID)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.GameID, &game.EventName, &game.SectionID, &game.Round, &game.Result,
			&game.Player1ID, &game.Player1Name, &game.Player1Color,
			&game.Player2ID, &game.Player2Name, &game.Player2Color); err != nil {
			utils.CheckErr(err)
			return games, err
		}
		games = append(games, game)
	}
	return
}

// GetPlayersWithRating fetches all players who ever had a post tournament rating from start to end - 1.
func (stats *StatsDB) GetPlayersInRatingRangeAndNoFide(start int, end int) (players []int, err error) {
	rows, err := stats.db.Query(queryPlayersInRatingRangeAndNoFide, start, end)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var player int
		if err := rows.Scan(&player); err != nil {
			utils.CheckErr(err)
			return players, err
		}
		players = append(players, player)
	}
	return
}
