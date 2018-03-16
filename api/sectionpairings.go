package main

import "uschess/statsdb"

// RoundResult is used to store result of each round for a player.
type RoundResult struct {
	PlayerID    int    `json:"playerId"`
	PlayerName  string `json:"playerName"`
	PlayerColor uint8  `json:"playerColor"`
	Result      string `json:"result,omitempty"`
}

type SectionPairings struct {
	SectionID string `json:"sectionId,omitempty"`
	// Map of PlayerId -> {Round -> RoundResult}.
	PlayerResults map[int]map[int]RoundResult `json:"player,omitempty"`
}

// NewSectionPairings returns a new initialized instance of SectionPairings.
func NewSectionPairings(sectionID string) SectionPairings {
	return SectionPairings{SectionID: sectionID, PlayerResults: make(map[int]map[int]RoundResult)}
}

func (sectionPairings *SectionPairings) addGames(games []statsdb.Game) {
	for _, game := range games {
		// We will create the map entries the first time we see the keys.
		if _, ok := sectionPairings.PlayerResults[game.Player1ID]; !ok {
			sectionPairings.PlayerResults[game.Player1ID] = make(map[int]RoundResult)
		}
		if _, ok := sectionPairings.PlayerResults[game.Player2ID]; !ok {
			sectionPairings.PlayerResults[game.Player2ID] = make(map[int]RoundResult)
		}
		roundResults := sectionPairings.gameToRoundResults(game)
		sectionPairings.addRoundInfo(game.Player1ID, game.Round, roundResults[1])
		sectionPairings.addRoundInfo(game.Player2ID, game.Round, roundResults[0])
	}
}

func (sectionPairings *SectionPairings) addRoundInfo(playerID int, round int, roundResult RoundResult) {
	// If the round information for the specified round is not in the map, then enter it and we are good.
	if _, ok := sectionPairings.PlayerResults[playerID][round]; !ok {
		sectionPairings.PlayerResults[playerID][round] = roundResult
		return
	}

	// If the round information for the specified round is already in the map, this is wrong. Most likely this is
	// due to this being a Round Robin event. In this case, there is no round information for this event at uscf.
	// As a result, we should just put this information in the first available round.

	round = 1 // There is no round 0. :)
	for {
		if _, ok := sectionPairings.PlayerResults[playerID][round]; !ok {
			sectionPairings.PlayerResults[playerID][round] = roundResult
			break
		}
		round++
	}
	return
}

// Convert game information into two results for each player. This will let UI clients
// to easily showcase performance against opponents easily.
func (sectionPairings *SectionPairings) gameToRoundResults(game statsdb.Game) (roundResults [2]RoundResult) {
	var player2Result string
	switch game.Result {
	case "W":
		player2Result = "L"
	case "L":
		player2Result = "W"
	case "D":
		player2Result = "D"
	}
	// roundResults[0] will contain information on Player2's result against Player1.
	roundResults[0] = RoundResult{game.Player1ID, game.Player1Name, game.Player1Color, player2Result}
	// roundResults[1] will contain information on Player1's result against Player2.
	roundResults[1] = RoundResult{game.Player2ID, game.Player2Name, game.Player2Color, game.Result}
	return roundResults
}
