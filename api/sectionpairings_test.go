package main

import (
	"reflect"
	"testing"
	"uschess/statsdb"
)

func TestAddGamesForRoundRobinEvent(t *testing.T) {
	var games = []statsdb.Game{
		{GameID: 653188, SectionID: "1.1", Round: 1, Result: "*", Player1ID: 13469010, Player1Name: "John Champion", Player1Color: 0, Player2ID: 0, Player2Name: "", Player2Color: 0},
		{GameID: 653189, SectionID: "1.1", Round: 1, Result: "L", Player1ID: 12448994, Player1Name: "Jeff Seconder", Player1Color: 1, Player2ID: 13469010, Player2Name: "John Champion", Player2Color: 0},
		{GameID: 653190, SectionID: "1.1", Round: 2, Result: "*", Player1ID: 12448994, Player1Name: "Jeff Seconder", Player1Color: 0, Player2ID: 0, Player2Name: "", Player2Color: 0},
		{GameID: 653191, SectionID: "1.1", Round: 1, Result: "L", Player1ID: 20048033, Player1Name: "Justin The Third", Player1Color: 2, Player2ID: 13469010, Player2Name: "John Champion", Player2Color: 0},
		{GameID: 653192, SectionID: "1.1", Round: 2, Result: "L", Player1ID: 20048033, Player1Name: "Justin The Third", Player1Color: 1, Player2ID: 12448994, Player2Name: "Jeff Seconder", Player2Color: 0},
		{GameID: 653193, SectionID: "1.1", Round: 3, Result: "*", Player1ID: 20048033, Player1Name: "Justin The Third", Player1Color: 0, Player2ID: 0, Player2Name: "", Player2Color: 0},
		{GameID: 653194, SectionID: "1.1", Round: 1, Result: "L", Player1ID: 16566837, Player1Name: "Logan Last", Player1Color: 1, Player2ID: 13469010, Player2Name: "John Champion", Player2Color: 0},
		{GameID: 653195, SectionID: "1.1", Round: 2, Result: "L", Player1ID: 16566837, Player1Name: "Logan Last", Player1Color: 1, Player2ID: 12448994, Player2Name: "Jeff Seconder", Player2Color: 0},
		{GameID: 653196, SectionID: "1.1", Round: 3, Result: "L", Player1ID: 16566837, Player1Name: "Logan Last", Player1Color: 2, Player2ID: 20048033, Player2Name: "Justin The Third", Player2Color: 0},
		{GameID: 653197, SectionID: "1.1", Round: 4, Result: "*", Player1ID: 16566837, Player1Name: "Logan Last", Player1Color: 0, Player2ID: 0, Player2Name: "", Player2Color: 0},
	}
	players := []int{13469010, 12448994, 20048033, 16566837}
	expectedPairings := SectionPairings{}
	expectedPairings.SectionID = "1:1"
	expectedPairings.PlayerResults = make(map[int]map[int]RoundResult)
	roundResults := []map[int]RoundResult{
		map[int]RoundResult{
			1: {13469010, "John Champion", 0, "L"},
			2: {12448994, "Jeff Seconder", 0, "L"},
			3: {20048033, "Justin The Third", 0, "L"},
			4: {0, "", 0, "*"},
		},
		map[int]RoundResult{
			1: {0, "", 0, "*"},
			2: {12448994, "Jeff Seconder", 1, "W"},
			3: {20048033, "Justin The Third", 2, "W"},
			4: {16566837, "Logan Last", 1, "W"},
		},
		map[int]RoundResult{
			1: {13469010, "John Champion", 0, "L"},
			2: {0, "", 0, "*"},
			3: {20048033, "Justin The Third", 1, "W"},
			4: {16566837, "Logan Last", 1, "W"},
		},
		map[int]RoundResult{
			1: {13469010, "John Champion", 0, "L"},
			2: {12448994, "Jeff Seconder", 0, "L"},
			3: {0, "", 0, "*"},
			4: {16566837, "Logan Last", 2, "W"},
		},
		map[int]RoundResult{
			1: {13469010, "John Champion", 0, ""},
			2: {12448994, "Jeff Seconder", 0, ""},
			3: {20048033, "Justin The Third", 0, ""},
			4: {16566837, "Logan Last", 0, ""},
		},
	}
	expectedPairings.PlayerResults[16566837] = roundResults[0]
	expectedPairings.PlayerResults[13469010] = roundResults[1]
	expectedPairings.PlayerResults[12448994] = roundResults[2]
	expectedPairings.PlayerResults[20048033] = roundResults[3]
	expectedPairings.PlayerResults[0] = roundResults[4]

	pairings := NewSectionPairings("1.1")
	pairings.addGames(games)
	for player := range players {
		if !reflect.DeepEqual(pairings.PlayerResults[player], expectedPairings.PlayerResults[player]) {
			t.Error("Parsing failed ", pairings, expectedPairings)
		}
	}
}
