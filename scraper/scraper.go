package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"uschess/statsdb"
	"uschess/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Parse Command line params.
	startDatePtr := flag.String("startdate", time.Now().Local().Format("2006-01-02"), "Start date to start scraping.")
	offsetPtr := flag.Int("offset", 0, "This is used as offset from the specified startdate. Could be negative.")
	monthPtr := flag.Bool("month", false, "Processes data for the rest of the month starting from startdate.")
	yearPtr := flag.Bool("year", false, "Processes data for the rest of the year starting from startdate.")
	forcePtr := flag.Bool("force", false, "Force update database. This deletes existing data from db and then writes the new data")
	savePtr := flag.Bool("save", false, "Saves data to database. By default save is false and only does the data fetch.")

	flag.Parse()

	t, err := time.Parse("2006-01-02", *startDatePtr)
	if err != nil {
		log.Fatalf("Invalid format for startdate specified. %s", err.Error())
	}
	utils.CheckErr(err)

	t = t.AddDate(0, 0, *offsetPtr)
	log.Printf("Starting date for processing: %s", t.String())

	// Initialize database connection.
	var stats statsdb.StatsDB
	stats.Open()
	defer stats.Close()
	log.Printf("Initialized database.")

	if *monthPtr {
		processMonth(stats, t, *forcePtr, *savePtr)
	} else if *yearPtr {
		processYear(stats, t, *forcePtr, *savePtr)
	} else {
		// Just the specified day
		processDate(stats, t, *forcePtr, *savePtr)
	}
}

// processYear does the same operation as processDate, but does it for the entire
// year starting at date t.
func processYear(stats statsdb.StatsDB, t time.Time, force bool, save bool) {
	curYear := t.Year()
	for t.Year() == curYear {
		processDate(stats, t, force, save)
		t = t.AddDate(0, 0, 1)
	}
}

// processMonth does the same operation as processDate, but does it for the entire
// month starting at date t.
func processMonth(stats statsdb.StatsDB, t time.Time, force bool, save bool) {
	curMonth := t.Month()
	for t.Month() == curMonth {
		processDate(stats, t, force, save)
		t = t.AddDate(0, 0, 1)
	}
}

// processDate for a given date t will fetch and write data into the database.
// if force is set, data for that date will be first deleted before writing nrew data.
// If save is set, data will be written to the database. If not, we just do a dry run without
// writing anything to the database.
func processDate(stats statsdb.StatsDB, t time.Time, force bool, save bool) {
	date := fmt.Sprintf("%4d-%02d-%02d", t.Year(), t.Month(), t.Day())
	log.Printf("Fetching all events for %s", date)
	events := fetchEvents(date)

	if force {
		if save {
			log.Printf("Deleting all events for date: %s", date)
			stats.DeleteEvents(date)
		} else {
			log.Printf("Dryrun deleting all events at date %s", date)
		}
	}

	for _, event := range events {
		if save {
			log.Printf("Saving event %s:%s", event.id, event.name)
			saveEvent(stats, event)
		} else {
			log.Printf("Dryrun saving event %s:%s", event.id, event.name)
		}
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
				stats.InsertTournamentHistory(entry.id, event.id, section.id, ratingChange.ratingType, preRating, postRating, entry.score)
			}
			uscfID[entry.position] = entry.id
		}
		for _, entry := range section.entries {
			for i, game := range entry.games {
				// The uscf rating tables store information for both players, we only need to store the game information once.
				// We are choosing to store when the second player's position is lower so as to accomodate for one player entries
				// such as H, X, F etc, when the second player position is considered 0.
				if game.player2 < game.player1 {
					stats.InsertGame(event.id, section.id, i+1, uscfID[game.player1], int(game.player1Color), uscfID[game.player2], int(game.player2Color), game.result)
				}
			}
		}
	}
}
