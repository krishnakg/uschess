package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"uschess/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

// Event is used to collect all information about an USChess event (tournament).
type Event struct {
	id          string
	numSections int
	state       string
	city        string
	date        string
	players     int
	name        string
	sections    map[string]Section
}

// Section is used to collect all information about a section in a USChess event.
type Section struct {
	id      string
	name    string
	entries []EventEntry
}

// EventEntry represents each row in the tournament cross-table.
type EventEntry struct {
	position int
	score    float64
	id       int
	name     string
	state    string
	change   []RatingChange
	games    []Game
}

// RatingChange represents the rating change for a player in the specified event type.
type RatingChange struct {
	ratingType string
	pre        string
	post       string
}

// Game represents the result of a chess game between two players with the color information if available.
type Game struct {
	result       string
	player1      int
	player1Color Color
	player2      int
	player2Color Color
}

// Color of the pieces
type Color int

const (
	// UNKNOWN means US Chess data did not have the color information of the players.
	UNKNOWN Color = iota
	WHITE
	BLACK
)

// Top level urls to Parse
const (
	eventSearchPage = "http://www.uschess.org/datapage/event-search.php"
	eventTablePage  = "http://www.uschess.org/msa/XtblMain.php"
)

// Fetch all events for a gioven date
func fetchEvents(date string) []Event {
	eventSearchURL := fmt.Sprintf("%s?name=&state=ANY&city=&date_from=%s&date_to=%s&order=D&minsize=&affil=&timectl=&mode=Find", eventSearchPage, date, date)
	doc, err := goquery.NewDocument(eventSearchURL)
	if err != nil {
		glog.Fatal(err)
	}
	events := make([]Event, 0)
	cols := make([]string, 7)
	doc.Find(".blog form tr").Each(func(i int, s *goquery.Selection) {
		lines := s.Find("td")
		if lines.Size() != 7 {
			return
		}

		lines.Each(func(j int, p *goquery.Selection) {
			cols[j] = strings.TrimSpace(p.Text())
		})

		numSections, err := strconv.Atoi(cols[1])
		if err != nil {
			// This fails for the header row. So this is not an error.
			return
		}
		players, err := strconv.Atoi(cols[5])
		if err != nil {
			glog.Fatalf("Unable to parse players in the line %s", cols[5])
			return
		}

		event := Event{
			id:          cols[0],
			numSections: numSections,
			state:       cols[2],
			city:        cols[3],
			date:        cols[4],
			players:     players,
			name:        cols[6],
			sections:    make(map[string]Section),
		}
		glog.Infof("Fetching event %s:%s", event.id, event.name)
		for _, section := range fetchResultTable(cols[0], numSections) {
			event.sections[section.id] = section
		}
		events = append(events, event)

	})
	return events

}

// Fetch information about all sections in this event.
func fetchResultTable(event string, numSections int) []Section {
	eventTableURL := fmt.Sprintf("%s?%s.0", eventTablePage, event)

	doc, err := goquery.NewDocument(eventTableURL)
	if err != nil {
		glog.Fatal(err)
	}

	sectionNames := make(map[int]string)
	doc.Find("td [bgcolor=\"DDDDFF\"] b").Each(func(j int, s *goquery.Selection) {
		section, name := parseSectionName(s.Text())
		sectionNames[section] = name
	})

	var sections []Section
	doc.Find("pre").Each(func(i int, s *goquery.Selection) {
		var section Section
		section.id = fmt.Sprintf("%s.%d", event, i+1)
		section.name = sectionNames[i+1]

		glog.Infof("Fetching section %s:%s", section.id, section.name)
		re := regexp.MustCompile("-+\n")
		array := re.Split(s.Text(), -1)

		for _, str := range array[3 : len(array)-1] {
			entry := parseSectionEntry(str)
			section.entries = append(section.entries, entry)
			glog.Infof("Processed player with id :%d", entry.id)
		}
		sections = append(sections, section)
	})
	return sections
}
func parseSectionEntry(entryStr string) EventEntry {
	var entry EventEntry
	var err error

	lines := strings.Split(entryStr, "\n")

	/* Process Line 1
	Col 0: Position of player
	Col 1: Name of player
	Col 2: Total Score of the player
	Col 3 onwards, pair of results: (Result OpponentPosition)
	*/
	parts := strings.Split(lines[0], "|")
	entry.position, err = strconv.Atoi(strings.TrimSpace(parts[0]))
	utils.CheckErr(err)
	entry.name = strings.Title(strings.ToLower(strings.TrimSpace(parts[1])))
	entry.score, err = strconv.ParseFloat(strings.TrimSpace(parts[2]), 32)
	utils.CheckErr(err)

	// '*' in the result field represents a round robin table where the player cannot play himself.
	re := regexp.MustCompile(` *([A-Z\*]) *([0-9]+)* *`)
	for _, part := range parts[3 : len(parts)-1] {
		result := re.FindStringSubmatch(part)
		if len(result) != 3 {
			// If the round information is empty, there is nothing to process for this round
			continue
		}
		// Many entries like F, U, B, X do not need an opponent. So we will assign that as default
		// and fetch the opponent only if it is present.
		player2 := 0
		if len(result[2]) > 0 {
			player2, err = strconv.Atoi(result[2])
			utils.CheckErr(err)
		}

		game := Game{result[1], entry.position, 0, player2, 0}
		entry.games = append(entry.games, game)
	}

	/* Process Line 2.
	Col 0: State
	Col 1: USCF Id / RatingType : Pre Rating -> Post Rating
	Col 2: Any norm information. Ignoring for now.
	Col 3 onwards, if present color information for the round.
	*/
	parts = strings.Split(lines[1], "|")
	// Col 0
	entry.state = strings.TrimSpace(parts[0])

	// Col 1
	re = regexp.MustCompile(` *([0-9]*)[ /]+([A-Z]+)*[: ]*([0-9A-Za-z]+)*[ \->]*([0-9Pp]*)`)
	result := re.FindStringSubmatch(parts[1])
	if result[1] != "" {
		entry.id, err = strconv.Atoi(result[1])
		utils.CheckErr(err)
	} else {
		entry.id = 0
	}
	ratingChange := RatingChange{result[2], result[3], result[4]}
	entry.change = append(entry.change, ratingChange)

	// Col 3+
	for i := 0; i < len(entry.games); i++ {
		colorStr := strings.TrimSpace(parts[3+i])
		if colorStr == "W" {
			entry.games[i].player1Color = WHITE
			entry.games[i].player2Color = BLACK
		} else if colorStr == "B" {
			entry.games[i].player1Color = BLACK
			entry.games[i].player2Color = WHITE
		}
	}

	// Process line 3 if present
	if len(lines) >= 3 {
		parts := strings.Split(lines[2], "|")
		if len(parts) >= 2 {
			re := regexp.MustCompile(` *([A-Z]+)[: ]+([0-9A-Za-z]+)[ \->]+([0-9Pp]+)`)
			result = re.FindStringSubmatch(parts[1])
			ratingChange := RatingChange{result[1], result[2], result[3]}
			entry.change = append(entry.change, ratingChange)
		}
	}

	return entry
}

func parseSectionName(sectionStr string) (int, string) {
	re := regexp.MustCompile(` *Section ([0-9]+) - (.*)`)
	parts := re.FindStringSubmatch(sectionStr)
	section, err := strconv.Atoi(parts[1])
	utils.CheckErr(err)
	return section, parts[2]
}

func parseRating(ratingStr string) (rating int, games int) {
	re := regexp.MustCompile(` *([0-9]+)[Pp]*([0-9]*)`)
	parts := re.FindStringSubmatch(ratingStr)
	if len(parts) < 2 {
		return
	}
	rating, err := strconv.Atoi(parts[1])
	utils.CheckErr(err)
	if parts[2] != "" {
		games, err = strconv.Atoi(parts[2])
		utils.CheckErr(err)
	}
	return rating, games
}
