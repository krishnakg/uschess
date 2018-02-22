package main

import (
	"reflect"
	"testing"
)

func TestParseSectionEntryWithTwoRatings(t *testing.T) {
	entryStr := "    1 | ANDY S PORTER                   |3.0  |W   6|W   4|W   3|\n" +
		"   IN | 12425190 / R: 2029   ->2037     |     |     |     |     |\n" +
		"      |            Q: 2003   ->2012     |     |     |     |     |\n"
	expectedEntry := EventEntry{1, 3.0, 12425190, "Andy S Porter", "IN",
		[]RatingChange{
			{"R", "2029", "2037"},
			{"Q", "2003", "2012"},
		}, []Game{
			{"W", 1, 0, 6, 0},
			{"W", 1, 0, 4, 0},
			{"W", 1, 0, 3, 0},
		}}
	testParseSectionEntry(t, entryStr, expectedEntry)
}

func TestParseSectionEntryWithOneRatingAndUnplayedRounds(t *testing.T) {
	entryStr := "   21 | DR BENJAMIN KARREN              |1.0  |W   8|L  18|F    |U    |U    |U    |U    |\n" +
		"   AZ | 13468860 / OB: 1572   ->1565    |     |B    |B    |W    |     |     |     |     |\n"
	expectedEntry := EventEntry{21, 1.0, 13468860, "Dr Benjamin Karren", "AZ",
		[]RatingChange{
			{"OB", "1572", "1565"},
		}, []Game{
			{"W", 21, 2, 8, 1},
			{"L", 21, 2, 18, 1},
			{"F", 21, 1, 0, 2},
			{"U", 21, 0, 0, 0},
			{"U", 21, 0, 0, 0},
			{"U", 21, 0, 0, 0},
			{"U", 21, 0, 0, 0},
		}}
	testParseSectionEntry(t, entryStr, expectedEntry)
}

func TestParseSectionEntryWithProvisionalRating(t *testing.T) {
	entryStr := "    8 | SHENGHAN XU                     |0.0  |L   3|L   7|L   5|L   6|L   4|\n" +
		"   VA | 16534371 / R:  109P7 -> 110P12  |     |     |     |     |     |     |\n" +
		"   |            Q:  109P7 -> 101P12  |     |     |     |     |     |     |\n"
	expectedEntry := EventEntry{8, 0.0, 16534371, "Shenghan Xu", "VA",
		[]RatingChange{
			{"R", "109P7", "110P12"},
			{"Q", "109P7", "101P12"},
		}, []Game{
			{"L", 8, 0, 3, 0},
			{"L", 8, 0, 7, 0},
			{"L", 8, 0, 5, 0},
			{"L", 8, 0, 6, 0},
			{"L", 8, 0, 4, 0},
		}}
	testParseSectionEntry(t, entryStr, expectedEntry)
}

func TestParseSectionEntryForRoundRobinEntry(t *testing.T) {
	entryStr := "    2 | THOMAS J BELKE                  |1.5  |W   4|*    |L   1|D   3|\n" +
		"   VA | 10126550 / Q: 1778   ->1792     |     |B    |     |W    |B    |\n"
	expectedEntry := EventEntry{2, 1.5, 10126550, "Thomas J Belke", "VA",
		[]RatingChange{
			{"Q", "1778", "1792"},
		}, []Game{
			{"W", 2, 2, 4, 1},
			{"*", 2, 0, 0, 0},
			{"L", 2, 1, 1, 2},
			{"D", 2, 2, 3, 1},
		}}
	testParseSectionEntry(t, entryStr, expectedEntry)
}

func TestParseSectionEntryForNoRating(t *testing.T) {
	entryStr := "   69 | WAYNE T FISCHER                 |0.0  |U    |U    |U    |U    |\n" +
		"   NJ | 12894588 /                      |     |     |     |     |     |\n"
	expectedEntry := EventEntry{69, 0.0, 12894588, "Wayne T Fischer", "NJ",
		[]RatingChange{
			{"", "", ""},
		}, []Game{
			{"U", 69, 0, 0, 0},
			{"U", 69, 0, 0, 0},
			{"U", 69, 0, 0, 0},
			{"U", 69, 0, 0, 0},
		}}
	testParseSectionEntry(t, entryStr, expectedEntry)
}

// Convinience function for the above tests.
func testParseSectionEntry(t *testing.T, entryStr string, expectedEntry EventEntry) {
	entry := parseSectionEntry(entryStr)
	// Need to check if DeepEqual is the right way to do this comparision.
	if !reflect.DeepEqual(entry, expectedEntry) {
		t.Error("Parsing failed ", entry, expectedEntry)
	}
}

func TestParseSectionName(t *testing.T) {
	section, name := parseSectionName(" Section 8 - QUADS 9 ")
	if section != 8 && name != "QUADS 9" {
		t.Error("Parsing section and name failed", section, name)
	}

	section, name = parseSectionName(" Section 18 - QUADS 19 ")
	if section != 18 && name != "QUADS 19" {
		t.Error("Parsing section and name failed", section, name)
	}
}

func TestParseRating(t *testing.T) {
	ratingStr := " 845P34 "
	rating, games := parseRating(ratingStr)
	if rating != 845 && games != 34 {
		t.Errorf("Parsing rating string %s gave %d %d", ratingStr, rating, games)
	}

	ratingStr = " 845"
	rating, games = parseRating(ratingStr)
	if rating != 845 && games != 0 {
		t.Errorf("Parsing rating string %s gave %d %d", ratingStr, rating, games)
	}

	rating, games = parseRating(" ")
	if rating != 0 && games != 0 {
		t.Errorf("Parsing empty rating string gave %d %d", rating, games)
	}
}
