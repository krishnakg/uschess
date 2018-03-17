package main

import "fmt"

// Top level urls to Parse
const (
	eventSearchPage = "http://www.uschess.org/datapage/event-search.php"
	eventTablePage  = "http://www.uschess.org/msa/XtblMain.php"
	playerPage      = "http://www.uschess.org/assets/msa_joomla/MbrDtlMain.php"
)

func getEventSearchURL(date string) string {
	return fmt.Sprintf("%s?name=&state=ANY&city=&date_from=%s&date_to=%s&order=D&minsize=&affil=&timectl=&mode=Find", eventSearchPage, date, date)
}

func getEventTableURL(eventID string) string {
	return fmt.Sprintf("%s?%s.0", eventTablePage, eventID)
}

func getPlayerPageURL(playerID int) string {
	return fmt.Sprintf("%s?%d", playerPage, playerID)
}
