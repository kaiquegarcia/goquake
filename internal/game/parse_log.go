package game

import (
	"goquake/internal/config"
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
		switch config.Get().GroupBy {
		case "mean":
			result = game.KillsByMeans()
		default:
			result = game.Stats(config.Get().ConsiderWorldAsPlayer)
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
