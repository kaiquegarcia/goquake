package game

import (
	"goquake/internal/config"
	"sort"
	"strconv"
	"strings"
)

var initGamePattern = "------------------------------------------------------------"

func ParseLog(log string) map[string]interface{} {
	games := make(map[string]interface{})
	gameLogs := strings.Split(log, initGamePattern)[1:]
	gameCounter := 1
	for index := 0; index < len(gameLogs); index++ {
		if len(gameLogs[index]) < 10 {
			continue
		}

		gameEvents := parseGameLog(gameLogs[index])
		game := New(gameEvents)

		var result interface{}
		if config.Get().Ranking {
			result = generateResultForRanking(game)
		} else {
			result = generateResultForGroupBy(game)
		}

		number := strconv.Itoa(gameCounter)
		if gameCounter < 10 {
			number = "0" + number
		}
		games["game-"+number] = result
		gameCounter++
	}

	return games
}

func generateResultForRanking(game Data) interface{} {
	rankedPlayers := make([]RankedPlayer, 0)

	kills := game.Stats(false).Kills
	if len(kills) == 0 {
		return rankedPlayers
	}

	for name, killCount := range kills {
		rankedPlayers = append(rankedPlayers, RankedPlayer{
			Name:  name,
			Kills: killCount,
		})
	}

	sort.SliceStable(rankedPlayers, func(i, j int) bool {
		return rankedPlayers[i].Kills > rankedPlayers[j].Kills
	})

	for index := 0; index < len(rankedPlayers); index++ {
		rankedPlayers[index].Rank = index + 1
	}

	return rankedPlayers
}

func generateResultForGroupBy(game Data) interface{} {
	switch config.Get().GroupBy {
	case "mean":
		return game.KillsByMeans()
	}

	return game.Stats(config.Get().ConsiderWorldAsPlayer)
}
