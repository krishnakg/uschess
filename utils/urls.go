package utils

import "fmt"

// Top level urls to Parse
const (
	eventSearchPage = "http://www.uschess.org/datapage/event-search.php"
	eventTablePage  = "http://www.uschess.org/msa/XtblMain.php"
	playerPage      = "http://www.uschess.org/assets/msa_joomla/MbrDtlMain.php"
)

// GetEventSearchURL constructs the url to fetch list of events for the given date.
func GetEventSearchURL(date string) string {
	return fmt.Sprintf("%s?name=&state=ANY&city=&date_from=%s&date_to=%s&order=D&minsize=&affil=&timectl=&mode=Find", eventSearchPage, date, date)
}

// GetEventTableURL returns the URL to the event's result table.
func GetEventTableURL(eventID string) string {
	return fmt.Sprintf("%s?%s.0", eventTablePage, eventID)
}

// GetPlayerPageURL returns the URL to the specified player's USCF page.
func GetPlayerPageURL(playerID int) string {
	return fmt.Sprintf("%s?%d", playerPage, playerID)
}
