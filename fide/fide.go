package main

import (
	"flag"
	"regexp"
	"strconv"
	"uschess/statsdb"
	"uschess/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"

	_ "github.com/go-sql-driver/mysql"
)

func fetchFideID(uscfID int) (fideID int, err error) {
	doc, err := goquery.NewDocument(utils.GetPlayerPageURL(uscfID))
	if err != nil {
		glog.Fatal(err)
	}

	re := regexp.MustCompile("idcode=([0-9]*)")

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		url, found := s.Attr("href")
		if found {
			result := re.FindStringSubmatch(url)
			if len(result) == 2 {
				fideID, err = strconv.Atoi(result[1])
			}
		}
	})

	return fideID, err
}

func main() {
	startPtr := flag.Int("start", 0, "Starting rating to look for fide id")
	endPtr := flag.Int("end", 0, "Ending rating to look for fide id")

	flag.Parse()
	defer glog.Flush()

	if *startPtr < 0 || *endPtr < 0 {
		glog.Fatal("Invalid values for start and end rating")
	}

	// Initialize database connection.
	var stats statsdb.StatsDB
	stats.Open()
	defer stats.Close()
	glog.Info("Initialized database.")

	players, _ := stats.GetPlayersInRatingRangeAndNoFide(*startPtr, *endPtr)
	glog.Infof("Number of players %d rated between %d and %d", len(players), *startPtr, *endPtr)
	for _, uscfID := range players {
		fideID, err := fetchFideID(uscfID)
		if err != nil {
			glog.Error(err)
			continue
		}
		if fideID != 0 {
			stats.InsertFideID(uscfID, fideID)
			glog.Infof("%d %d", uscfID, fideID)
		}
	}
}
